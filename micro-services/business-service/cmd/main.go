package main

import (
	"fmt"

	"github.com/ronexlemon/rail/micro-services/business-service/events"
)

func main(){
	fmt.Println("BUSINESS SERVICE .......")
	events.ConsumeRegister()
	//select{}
}