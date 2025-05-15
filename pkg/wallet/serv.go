package wallet

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
)

const (
	// SERV address prefix
	ServPrefix = "serv"
)

// ServWallet represents a SERV wallet
type ServWallet struct {
	Address    string
	PrivateKey *secp256k1.PrivKey
	PublicKey  *secp256k1.PubKey
}

// NewServWallet creates a new SERV wallet from a private key
func NewServWallet(privKeyHex string) (*ServWallet, error) {
	// Decode private key from hex
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key hex: %w", err)
	}

	privKey := &secp256k1.PrivKey{Key: privKeyBytes}
	pubKey := privKey.PubKey().(*secp256k1.PubKey)

	// Generate SERV address
	addr, err := ConvertToServAddress(sdk.AccAddress(pubKey.Address()))
	if err != nil {
		return nil, fmt.Errorf("failed to create SERV address: %w", err)
	}

	return &ServWallet{
		Address:    addr,
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}, nil
}

// CreateNewServWallet creates a completely new SERV wallet
func CreateNewServWallet() (*ServWallet, error) {
	// Generate new private key
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey().(*secp256k1.PubKey)

	// Generate SERV address
	addr, err := ConvertToServAddress(sdk.AccAddress(pubKey.Address()))
	if err != nil {
		return nil, fmt.Errorf("failed to create SERV address: %w", err)
	}

	return &ServWallet{
		Address:    addr,
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}, nil
}

// ConvertToServAddress converts a Cosmos SDK address to a SERV address
func ConvertToServAddress(addr sdk.AccAddress) (string, error) {
	bech32Addr, err := bech32.ConvertAndEncode(ServPrefix, addr.Bytes())
	if err != nil {
		return "", fmt.Errorf("failed to encode SERV address: %w", err)
	}
	return bech32Addr, nil
}

// ConvertFromServAddress converts a SERV address to a Cosmos SDK address
func ConvertFromServAddress(servAddr string) (sdk.AccAddress, error) {
	if !strings.HasPrefix(servAddr, ServPrefix) {
		return nil, fmt.Errorf("invalid SERV address prefix")
	}

	_, bz, err := bech32.DecodeAndConvert(servAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode SERV address: %w", err)
	}

	return sdk.AccAddress(bz), nil
}

// Sign signs a message with the wallet's private key
func (w *ServWallet) Sign(msg []byte) ([]byte, error) {
	if w.PrivateKey == nil {
		return nil, fmt.Errorf("wallet has no private key")
	}
	return w.PrivateKey.Sign(msg)
}

// GetAddress returns the wallet's address
func (w *ServWallet) GetAddress() string {
	return w.Address
}

// ExportPrivateKey exports the private key as a hex string
func (w *ServWallet) ExportPrivateKey() string {
	return hex.EncodeToString(w.PrivateKey.Key)
}

// ValidateServAddress validates a SERV address format
func ValidateServAddress(address string) error {
	if !strings.HasPrefix(address, ServPrefix) {
		return fmt.Errorf("invalid SERV address prefix")
	}

	_, _, err := bech32.DecodeAndConvert(address)
	if err != nil {
		return fmt.Errorf("invalid SERV address format: %w", err)
	}

	return nil
}
