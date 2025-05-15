package blockchain

import (
	"context"
	"fmt"
	"os"

	"github.com/nutcase/dao-golang/pkg/auth"
	"github.com/nutcase/dao-golang/pkg/wallet"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signing "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
)

// Client represents a blockchain client
type Client struct {
	ctx          context.Context
	clientCtx    client.Context
	grpcConn     *grpc.ClientConn
	keyring      keyring.Keyring
	chainID      string
	defaultDenom string
	roleManager  *auth.RoleManager
}

// NewClient creates a new blockchain client
func NewClient(nodeURL string) (*Client, error) {
	// Initialize context
	ctx := context.Background()

	// Initialize codec
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	// Initialize keyring
	kr, err := keyring.New(
		"taskbounty",
		keyring.BackendMemory,
		"~/.taskbounty",
		os.Stdin,
		marshaler,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize keyring: %w", err)
	}

	// Connect to gRPC
	grpcConn, err := grpc.Dial(
		nodeURL,
		grpc.WithInsecure(), // Note: In production, use secure connection
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC: %w", err)
	}

	// Initialize client context
	clientCtx := client.Context{}
	clientCtx = clientCtx.WithKeyring(kr)
	clientCtx = clientCtx.WithNodeURI(nodeURL)
	clientCtx = clientCtx.WithChainID("taskbounty-testnet")
	clientCtx = clientCtx.WithBroadcastMode("block")
	clientCtx = clientCtx.WithGRPCClient(grpcConn)

	// Initialize role manager
	roleManager := auth.NewRoleManager()

	return &Client{
		ctx:          ctx,
		clientCtx:    clientCtx,
		grpcConn:     grpcConn,
		keyring:      kr,
		chainID:      "taskbounty-testnet",
		defaultDenom: "stake",
		roleManager:  roleManager,
	}, nil
}

// Close closes the client connections
func (c *Client) Close() error {
	if c.grpcConn != nil {
		return c.grpcConn.Close()
	}
	return nil
}

// TransferTokens transfers tokens from one address to another
func (c *Client) TransferTokens(from, to string, amount uint64) error {
	// Convert addresses
	fromAddr, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		return fmt.Errorf("invalid from address: %w", err)
	}

	toAddr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		return fmt.Errorf("invalid to address: %w", err)
	}

	// Create transfer message
	msg := banktypes.NewMsgSend(
		fromAddr,
		toAddr,
		sdk.NewCoins(sdk.NewInt64Coin(c.defaultDenom, int64(amount))),
	)

	// Create transaction factory
	txf := tx.Factory{}.WithChainID(c.chainID).WithTxConfig(c.clientCtx.TxConfig)

	// Build transaction
	txBuilder, err := txf.BuildUnsignedTx(msg)
	if err != nil {
		return fmt.Errorf("failed to build transaction: %w", err)
	}

	// Get account info
	accountClient := authtypes.NewQueryClient(c.grpcConn)
	accountRes, err := accountClient.Account(c.ctx, &authtypes.QueryAccountRequest{Address: fromAddr.String()})
	if err != nil {
		return fmt.Errorf("failed to get account info: %w", err)
	}

	// Get account info
	var account authtypes.AccountI
	err = c.clientCtx.Codec.UnpackAny(accountRes.Account, &account)
	if err != nil {
		return fmt.Errorf("failed to unpack account: %w", err)
	}

	sigData := signing.SingleSignatureData{
		SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}
	sig := signing.SignatureV2{
		PubKey:   account.GetPubKey(),
		Data:     &sigData,
		Sequence: account.GetSequence(),
	}

	err = txBuilder.SetSignatures(sig)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Broadcast transaction
	txBytes, err := c.clientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return fmt.Errorf("failed to encode transaction: %w", err)
	}

	res, err := c.clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return fmt.Errorf("failed to broadcast transaction: %w", err)
	}

	if res.Code != 0 {
		return fmt.Errorf("transaction failed: %s", res.RawLog)
	}

	return nil
}

// GetBalance gets the token balance for an address
func (c *Client) GetBalance(address string) (uint64, error) {
	// Convert address
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return 0, fmt.Errorf("invalid address: %w", err)
	}

	// Query balance
	bankClient := banktypes.NewQueryClient(c.grpcConn)
	res, err := bankClient.Balance(c.ctx, &banktypes.QueryBalanceRequest{
		Address: addr.String(),
		Denom:   c.defaultDenom,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to query balance: %w", err)
	}

	return res.Balance.Amount.Uint64(), nil
}

// CreateServWallet creates a new SERV wallet
func (c *Client) CreateServWallet() (string, error) {
	// Create key info
	privKey := secp256k1.GenPrivKey()
	address := sdk.AccAddress(privKey.PubKey().Address())

	kr, err := keyring.New(
		"taskbounty",
		keyring.BackendMemory,
		c.clientCtx.KeyringDir,
		c.clientCtx.Input,
		c.clientCtx.Codec,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create keyring: %w", err)
	}

	// Store the key info
	_, err = kr.SaveOfflineKey(
		"taskbounty",
		privKey.PubKey(),
	)
	if err != nil {
		return "", fmt.Errorf("failed to save key info: %w", err)
	}

	// Assign default member role
	err = c.roleManager.AssignRole(address.String(), auth.RoleMember)
	if err != nil {
		return "", fmt.Errorf("failed to assign role: %w", err)
	}

	return address.String(), nil
}

// ImportServWallet imports an existing SERV wallet from private key
func (c *Client) ImportServWallet(privKeyHex string) (string, error) {
	// Import SERV wallet
	privKey := secp256k1.GenPrivKey()
	address := sdk.AccAddress(privKey.PubKey().Address())

	// Create key info
	kr, err := keyring.New(
		"taskbounty",
		keyring.BackendMemory,
		c.clientCtx.KeyringDir,
		c.clientCtx.Input,
		c.clientCtx.Codec,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create keyring: %w", err)
	}

	// Store the key info
	_, err = kr.SaveOfflineKey(
		"taskbounty",
		privKey.PubKey(),
	)
	if err != nil {
		return "", fmt.Errorf("failed to save key info: %w", err)
	}

	// Assign default member role
	err = c.roleManager.AssignRole(address.String(), auth.RoleMember)
	if err != nil {
		return "", fmt.Errorf("failed to assign role: %w", err)
	}

	return address.String(), nil
}

// ValidateAddress validates a SERV address
func (c *Client) ValidateAddress(address string) error {
	return wallet.ValidateServAddress(address)
}

// HasRole checks if an address has a specific role
func (c *Client) HasRole(address string, role auth.Role) bool {
	return c.roleManager.HasRole(address, role)
}

// AssignRole assigns a role to an address
func (c *Client) AssignRole(address string, role auth.Role) error {
	return c.roleManager.AssignRole(address, role)
}
