package client

import (
	"context"
	"fmt"
	// "os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Client represents a client for interacting with the Know Freedom testnet
type Client struct {
	ctx         client.Context
	grpcConn    *grpc.ClientConn
	chainID     string
	keyring     keyring.Keyring
	gasPrice    sdk.DecCoin
	gasAdjust   float64
}

// NewClient creates a new client for interacting with the Know Freedom testnet
func NewClient(networkConfigPath string) (*Client, error) {
	// Load network configuration
	viper.SetConfigFile(networkConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Initialize codec
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	// Initialize client context
	clientCtx := client.Context{}
	clientCtx = clientCtx.WithCodec(marshaler)
	
	// Set up network parameters
	chainID := viper.GetString("network.chain_id")
	rpcEndpoint := viper.GetString("network.rpc_endpoint")
	grpcEndpoint := viper.GetString("network.grpc_endpoint")

	// Connect to gRPC
	grpcConn, err := grpc.Dial(
		grpcEndpoint,
		grpc.WithInsecure(), // Note: In production, use secure connection
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC: %w", err)
	}

	// Parse gas price
	gasPriceStr := viper.GetString("gas.gas_price")
	gasPrice, err := sdk.ParseDecCoin(gasPriceStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse gas price: %w", err)
	}

	// Initialize keyring
	kr, err := keyring.New(
		"taskbounty",
		keyring.BackendTest, // Note: In production, use a more secure backend
		"~/.taskbounty",
		nil,
		marshaler, // Provide the codec
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize keyring: %w", err)
	}

	// Set up client context
	clientCtx = clientCtx.
		WithChainID(chainID).
		WithNodeURI(rpcEndpoint).
		WithKeyring(kr).
		WithBroadcastMode("sync")

	return &Client{
		ctx:         clientCtx,
		grpcConn:    grpcConn,
		chainID:     chainID,
		keyring:     kr,
		gasPrice:    gasPrice,
		gasAdjust:   viper.GetFloat64("gas.gas_adjustment"),
	}, nil
}

// Close closes the client connections
func (c *Client) Close() error {
	if c.grpcConn != nil {
		return c.grpcConn.Close()
	}
	return nil
}

// GetContext returns the client context
func (c *Client) GetContext() client.Context {
	return c.ctx
}

// SignAndBroadcast signs and broadcasts a transaction
func (c *Client) SignAndBroadcast(msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	// Build transaction
	txBuilder := c.ctx.TxConfig.NewTxBuilder()
	err := txBuilder.SetMsgs(msgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to set msgs: %w", err)
	}

	// Set gas and fee
	txBuilder.SetGasLimit(flags.DefaultGasLimit)
	// Convert DecCoin to Coin for fee
	fee := sdk.NewCoin(c.gasPrice.Denom, c.gasPrice.Amount.TruncateInt())
	txBuilder.SetFeeAmount(sdk.NewCoins(fee))

	// Get signer info
	info, err := c.keyring.Key(c.ctx.GetFromName())
	if err != nil {
		return nil, fmt.Errorf("failed to get key info: %w", err)
	}

	// Get account details
	addr, err := info.GetAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to get address: %w", err)
	}

	acc, err := c.ctx.AccountRetriever.GetAccount(c.ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	// Get public key
	pubKey, err := info.GetPubKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get public key: %w", err)
	}

	// Sign the transaction
	sig := signing.SignatureV2{
		PubKey: pubKey,
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
			Signature: nil,
		},
		Sequence: acc.GetSequence(),
	}

	err = txBuilder.SetSignatures(sig)
	if err != nil {
		return nil, fmt.Errorf("failed to set signatures: %w", err)
	}

	// Encode the transaction
	txBytes, err := c.ctx.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, fmt.Errorf("failed to encode tx: %w", err)
	}

	// Broadcast the transaction
	res, err := c.ctx.BroadcastTx(txBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to broadcast tx: %w", err)
	}

	return res, nil
}

// CreateWallet creates a new wallet
func (c *Client) CreateWallet(name string) (string, error) {
	record, _, err := c.keyring.NewMnemonic(
		name,
		keyring.English,
		sdk.FullFundraiserPath,
		keyring.DefaultBIP39Passphrase,
		hd.Secp256k1,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create wallet: %w", err)
	}

	addr, err := record.GetAddress()
	if err != nil {
		return "", fmt.Errorf("failed to get address: %w", err)
	}

	return addr.String(), nil
}

// GetBalance gets the token balance for an address
func (c *Client) GetBalance(address string, denom string) (sdk.Coin, error) {
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.Coin{}, fmt.Errorf("invalid address: %w", err)
	}

	queryClient := banktypes.NewQueryClient(c.grpcConn)
	res, err := queryClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{
			Address: addr.String(),
			Denom:   denom,
		},
	)
	if err != nil {
		return sdk.Coin{}, fmt.Errorf("failed to query balance: %w", err)
	}

	return *res.Balance, nil
}
