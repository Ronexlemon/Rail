package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type APIKeys struct{
	PublicKey string
	SecretKey string
}

func GenerateAPIKeys()(APIKeys,error){
	pubBytes := make([]byte,16)
	secBytes := make([]byte,32)
	if _, err := rand.Read(pubBytes); err !=nil{
		return APIKeys{PublicKey: "",SecretKey: ""}, err
	}

	if _, err := rand.Read(secBytes); err !=nil{
		return APIKeys{PublicKey: "",SecretKey: ""}, err
	}

	publicKey := fmt.Sprintf("api_%s",hex.EncodeToString(pubBytes))
	secretKey := fmt.Sprintf("sec_%s",hex.EncodeToString(secBytes))
	return APIKeys{PublicKey: publicKey,SecretKey: secretKey}, nil
}

