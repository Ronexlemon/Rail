package main

import (
	"fmt"

	"github.com/ronexlemon/rail/shared/kafka/topics"
)

func main(){
	fmt.Println("start registering")
	topics.RegisterTopics()
	fmt.Println("end registering")
}
