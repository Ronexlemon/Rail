package chains

import (
	"crypto/ecdsa"

	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)


func CreateEVMWallet() (common.Address, string) {
	generatedPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(generatedPrivateKey)
	privateKey := hexutil.Encode(privateKeyBytes)[2:]
	//fmt.Println(privateKey)
	publicKey := generatedPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, privateKey
}