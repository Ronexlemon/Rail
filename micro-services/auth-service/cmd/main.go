package main

import (
	"fmt"
	"log"
	
	"net/http"

	config "github.com/ronexlemon/rail/micro-services/auth-service/configs"
	"github.com/ronexlemon/rail/micro-services/auth-service/database"
	grpcserver "github.com/ronexlemon/rail/micro-services/auth-service/grserver"
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/handler"
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/service"

)

func main() {
	fmt.Println("🚀 Starting AUTH SERVICE .......")

	// 1. Connect to DB
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	fmt.Println("✅ Database connected")
	defer db.Client.Disconnect()

	// 2. Setup repository, service, and handler
	repo := repository.NewBusinessRepository()
	svc := service.NewBusinessService(repo)
	h := handler.NewBusinessHandler(svc)

	// 3. Get environment ports
	portHTTP := config.GetEnv("SERVICE_PORT", "8081")
	

	// 4. Start gRPC server in a goroutine
	go grpcserver.ServerGrpc()

	// 5. Start HTTP routes
	http.HandleFunc("/register-business", h.RegisterBusinessHandler)

	log.Printf("🌍 HTTP Auth service running on :%s", portHTTP)
	if err := http.ListenAndServe(":"+portHTTP, nil); err != nil {
		log.Fatalf(" Failed to start HTTP server: %v", err)
	}
}
