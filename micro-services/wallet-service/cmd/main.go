package main

import (
	"fmt"

	"github.com/ronexlemon/rail/micro-services/wallet-service/events"
)

func main(){
	fmt.Println("Wallet SERVICE .......")
	events.ConsumeRegister()
	//select{}
}