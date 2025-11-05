package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type APIKeys struct{
	publicKey string
	secretKey string
}

func GenerateAPIKeys(businessName string)(APIKeys,error){
	pubBytes := make([]byte,16)
	secBytes := make([]byte,32)
	if _, err := rand.Read(pubBytes); err !=nil{
		return APIKeys{publicKey: "",secretKey: ""}, err
	}

	if _, err := rand.Read(secBytes); err !=nil{
		return APIKeys{publicKey: "",secretKey: ""}, err
	}

	publicKey := fmt.Sprintf("pub_%s_%s",businessName,hex.EncodeToString(pubBytes))
	secretKey := fmt.Sprintf("sec_%s",hex.EncodeToString(secBytes))
	return APIKeys{publicKey: publicKey,secretKey: secretKey}, nil
}

