package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/ronexlemon/rail/micro-services/auth-service/configs"
	"github.com/ronexlemon/rail/micro-services/auth-service/database"
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/handler"
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/service"
	//"github.com/ronexlemon/rail/micro-services/auth-service/utils"
)

func main(){
	fmt.Println("AUTH SERVICE .......")
	db,err := database.ConnectDB()
	if err !=nil{
		log.Fatal("Cannot connect to db",err)

	}
	fmt.Println("DB runing")
	defer db.Client.Disconnect()
	repo:= repository.NewBusinessRepository()
	svc := service.NewBusinessService(repo)
	h := handler.NewBusinessHandler(svc)
	port := config.GetEnv("SERVICE_PORT","8081")

	http.HandleFunc("/register-business",h.RegisterBusinessHandler)
	log.Printf("Auth service running on :%s",port)
	log.Fatal(http.ListenAndServe(":"+port,nil))


}