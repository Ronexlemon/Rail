package main

import (
	"fmt"

	"github.com/ronexlemon/rail/micro-services/transaction-service/events"
)

func main(){
	fmt.Println("TRANSACTION SERVICE .......")
	events.ConsumeRegister()
	//select{}
}