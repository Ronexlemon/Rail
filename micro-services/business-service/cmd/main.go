package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ronexlemon/rail/micro-services/business-service/configs"
	"github.com/ronexlemon/rail/micro-services/business-service/database"
	"github.com/ronexlemon/rail/micro-services/business-service/events"
	"github.com/ronexlemon/rail/micro-services/business-service/internal/handler"
	"github.com/ronexlemon/rail/micro-services/business-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/business-service/internal/service"
)

func main(){
	fmt.Println("BUSINESS SERVICE .......")
	db,err := database.ConnectDB()
	if err !=nil{
		log.Fatal("Cannot connect to db",err)

	}
	fmt.Println("DB runing")
	defer db.Client.Disconnect()
	go events.ConsumeRegister()
	repo:= repository.NewBusinessRepository()
	service:= service.NewBusinesssService(repo)
	h:= handler.NewBusinessHandlerService(service)
	port := configs.GetEnv("SERVICE_PORT","8085")
	http.HandleFunc("/register-business",h.RegisterBusinessHandler)
	log.Printf("Auth service running on :%s",port)
	log.Fatal(http.ListenAndServe(":"+port,nil))
	//select{}
}