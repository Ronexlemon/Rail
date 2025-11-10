package chains

import (
	//"encoding/base58"
	//"fmt"

	// Corrected imports after the library change
	"github.com/blocto/solana-go-sdk/types"
	"github.com/mr-tron/base58"
)

// WalletPair holds the public address and private key of the new wallet.
type WalletPair struct {
	Address    string 
	PrivateKey string 
}

// CreateSolanaWallet generates a new Ed25519 keypair and returns the 
// public address and the private key, both encoded in Base58.
func CreateSolanaWallet() (WalletPair, error) {
	account := types.NewAccount()

	address := account.PublicKey.ToBase58()

	privateKeyBytes := account.PrivateKey
	privateKeyBase58 := base58.Encode(privateKeyBytes)

	return WalletPair{
		Address:    address,
		PrivateKey: privateKeyBase58,
	}, nil
}