package main

import (
	"fmt"
	"log"

	//"github.com/ronexlemon/rail/micro-services/wallet-service/utils"
	//"github.com/ronexlemon/rail/micro-services/wallet-service/utils/chains"
	"github.com/ronexlemon/rail/micro-services/wallet-service/database"
	"github.com/ronexlemon/rail/micro-services/wallet-service/events"
	"github.com/ronexlemon/rail/micro-services/wallet-service/grpcserver"
)

func main(){
	fmt.Println("Wallet SERVICE .......")
	db,err := database.ConnectDB()
	if err !=nil{
		log.Fatal("Cannot connect to db",err)

	}
	fmt.Println("DB runing")
	defer func() {
	if err := db.Client.Disconnect(); err != nil {
		log.Printf("Error disconnecting DB: %v", err)
	}
}()
	go events.ConsumeRegister()
	go grpcserver.ServerGrpc()
	
	select{}
}