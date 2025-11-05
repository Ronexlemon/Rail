package main

import (
	"fmt"

	"github.com/ronexlemon/rail/micro-services/auth-service/utils"
)

func main(){
	fmt.Println("AUTH SERVICE .......")
	result,_:= utils.GenerateAPIKeys("RAILS")
	fmt.Println(result)

}