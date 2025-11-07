package main

import (
	"fmt"

	"github.com/ronexlemon/rail/micro-services/notification-service/events"
)

func main(){
	fmt.Println("Notification Service .......")
	events.ConsumeRegister()
	select{}
}