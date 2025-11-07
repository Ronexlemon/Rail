package main

import (
	"fmt"

	"github.com/ronexlemon/rail/micro-services/settlement-service/events"
)

func main(){
	fmt.Println("SETTLEMENT SERVICE .......")
	events.ConsumeRegister()
	select{}
}