package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ronexlemon/rail/micro-services/wallet-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/wallet-service/internal/service"
	"github.com/ronexlemon/rail/micro-services/wallet-service/prisma/db"
	pb "github.com/ronexlemon/rail/micro-services/wallet-service/proto"
	"google.golang.org/grpc"
)


type WalletGRPCServer struct{
	pb.UnimplementedWalletServiceServer
	repo *repository.WalletRepository
	service *service.WalletService
}

func NewWalletGrpcServer(repo *repository.WalletRepository, service *service.WalletService)(*WalletGRPCServer){
	return &WalletGRPCServer{
		repo: repo,
		service: service,
	}
}

func (s *WalletGRPCServer) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.CreateWalletResponse, error) {
	
	if req.BusinessId == "" && req.CustomerId == "" {
		return nil, fmt.Errorf("either business_id or customer_id must be provided")
	}

	
	var customerID *string
	if req.CustomerId != "" {
		customerID = &req.CustomerId
	}

	
	var walletType db.WalletType
	if req.BusinessId != "" && req.CustomerId != "" {
		walletType = db.WalletTypeCustomer
	} else if req.CustomerId != "" {
		walletType = db.WalletTypeCustomer
	} else {
		walletType = db.WalletTypeBusiness
	}

	
	wallet, err := s.service.CreateWallet(ctx, req.BusinessId, customerID, walletType)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %v", err)
	}

	// Get values from Prisma getters
	var businessID, custID string
	if val, ok := wallet.BusinessID(); ok {
		businessID = val
	}
	if val, ok := wallet.CustomerID(); ok {
		custID = val
	}

	// Map DB enum to Proto enum
	var pbWalletType pb.WalletType
	switch walletType {
	case db.WalletTypeBusiness:
		pbWalletType = pb.WalletType_BUSINESS
	case db.WalletTypeCustomer:
		pbWalletType = pb.WalletType_CUSTOMER
	default:
		pbWalletType = pb.WalletType_BUSINESS
	}

	
	return &pb.CreateWalletResponse{
		WalletId:   wallet.ID,
		BusinessId: businessID,
		CustomerId: custID,
		Type:       pbWalletType,
		Message:    "Wallet created successfully",
	}, nil
}


func ServerGrpc() {
	repo := repository.NewWalletRepository()
	service := service.NewWalletService(repo)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWalletServiceServer(s, NewWalletGrpcServer(repo,service))
	log.Println("Business service running on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
