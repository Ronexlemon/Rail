package main

import (
	"fmt"
	"log"

	//"github.com/ronexlemon/rail/micro-services/wallet-service/utils"
	//"github.com/ronexlemon/rail/micro-services/wallet-service/utils/chains"
	"github.com/ronexlemon/rail/micro-services/wallet-service/database"
	"github.com/ronexlemon/rail/micro-services/wallet-service/events"
)

func main(){
	fmt.Println("Wallet SERVICE .......")
	db,err := database.ConnectDB()
	if err !=nil{
		log.Fatal("Cannot connect to db",err)

	}
	fmt.Println("DB runing")
	defer db.Client.Disconnect()
	events.ConsumeRegister()
	// address,privateKey:=chains.CreateEVMWallet()
	// fmt.Println("Address",address)
	// fmt.Println("PrivateKey",privateKey)
	// d,_:=chains.CreateSolanaWallet()
	// fmt.Println("Solana",d.Address)
	// fmt.Println("Solana priv",d.PrivateKey)
	// t,_:=chains.CreateTronWallet()
	// fmt.Println("Tron",t.Address)
	// fmt.Println("Tron",t.PrivateKey)
	//select{}
}