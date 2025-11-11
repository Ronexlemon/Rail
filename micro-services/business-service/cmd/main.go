package main

import (
	"fmt"
	"log"
	//"net"
	"net/http"

	"github.com/ronexlemon/rail/micro-services/business-service/configs"
	"github.com/ronexlemon/rail/micro-services/business-service/database"
	"github.com/ronexlemon/rail/micro-services/business-service/events"
	grserver "github.com/ronexlemon/rail/micro-services/business-service/grserver"
	"github.com/ronexlemon/rail/micro-services/business-service/internal/handler"
	"github.com/ronexlemon/rail/micro-services/business-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/business-service/internal/service"
	// "github.com/ronexlemon/rail/micro-services/business-service/proto"
	// "google.golang.org/grpc"
)

func main(){
	// fmt.Println("BUSINESS SERVICE .......")
	// db,err := database.ConnectDB()
	// if err !=nil{
	// 	log.Fatal("Cannot connect to db",err)

	// }
	// fmt.Println("DB runing")
	// defer db.Client.Disconnect()
	// go events.ConsumeRegister()
	// lis, err := net.Listen("tcp", ":8085")
	// s := grpc.NewServer()
	// proto.RegisterBusinessServiceServer(s, &grpcserver.Server{})
	// repo:= repository.NewBusinessRepository()
	// service:= service.NewBusinesssService(repo)
	// h:= handler.NewBusinessHandlerService(service)
	// port := configs.GetEnv("SERVICE_PORT","8085")
	// http.HandleFunc("/register-business",h.RegisterBusinessHandler)
	// log.Printf("Auth service running on :%s",port)
	// log.Fatal(http.ListenAndServe(":"+port,nil))
	// //select{}

	fmt.Println("BUSINESS SERVICE .......")

	// Connect to DB
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	defer db.Client.Disconnect()
	fmt.Println("DB running")

	// Start Kafka consumer in background
	go events.ConsumeRegister()

	// Initialize repository, service, handler
	repo := repository.NewBusinessRepository()
	svc := service.NewBusinesssService(repo)
	h := handler.NewBusinessHandlerService(svc)

	// ---------------------------
	// Start HTTP server
	// ---------------------------
	port := configs.GetEnv("SERVICE_PORT", "8085")
	go func() {
		http.HandleFunc("/register-business", h.RegisterBusinessHandler)
		log.Printf("HTTP service running on :%s", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal("HTTP server error:", err)
		}
	}()

	// ---------------------------
	// Start gRPC server
	// ---------------------------
	grserver.GrpServer()
	// Keep main alive
	select {}
}