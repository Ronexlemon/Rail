package chains

import (
	"crypto/ecdsa" // Must import the standard ecdsa package for the type assertion
	"crypto/sha256"
	"fmt"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)



func CreateTronWallet() (WalletPair, error) {
	
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return WalletPair{}, err
	}

	
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return WalletPair{}, fmt.Errorf("error casting public key to *ecdsa.PublicKey")
	}

	
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	
	// Skip the '04' prefix and hash the remaining bytes
	pubKeyHash := sha256.Sum256(publicKeyBytes[1:])
	ripemdHasher := crypto.NewKeccakState() // Using Keccak/SHA3 for hashing
	ripemdHasher.Write(pubKeyHash[:])
	finalHash := ripemdHasher.Sum(nil)
	finalHash = finalHash[12:] // Takes the last 20 bytes (Ethereum-style address hash)

	// Step 3: Add TRON Address Prefix (0x41)
	versionedPayload := append([]byte{0x41}, finalHash...)

	// Step 4: Calculate Checksum (Double SHA256 of the versioned payload)
	hash1 := sha256.Sum256(versionedPayload)
	hash2 := sha256.Sum256(hash1[:])
	checksum := hash2[:4] // Use the first 4 bytes as the checksum

	
	fullPayload := append(versionedPayload, checksum...)
	tronAddressBase58 := base58.Encode(fullPayload)

	return WalletPair{
		Address:    tronAddressBase58,
		PrivateKey: privateKeyHex,
	}, nil
}