package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/ronexlemon/rail/micro-services/wallet-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/wallet-service/internal/service"
	"github.com/ronexlemon/rail/micro-services/wallet-service/prisma/db"
	pb "github.com/ronexlemon/rail/micro-services/wallet-service/proto"

	//"github.com/ronexlemon/rail/micro-services/wallet-service/utils/helpers"
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
	// var businessID, custID string
	var  custID string
	// if val, ok := wallet.BusinessID(); ok {
	// 	businessID = val
	// }
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
		BusinessId:  wallet.BusinessID,
		CustomerId: custID,
		Type:       pbWalletType,
		Message:    "Wallet created successfully",
	}, nil
}

func (s *WalletGRPCServer) BusinessWallet(ctx context.Context, req *pb.BusinessWalletRequest) (*pb.BusinessWalletResponse, error) {
	if req.BusinessId == "" && req.CustomerId == "" {
		return nil, fmt.Errorf("either business_id or customer_id must be provided")
	}

	var customerID *string
	if req.CustomerId != "" {
		customerID = &req.CustomerId
	}

	// Fetch wallets from service
	wallets, err := s.service.BusinessWallet(ctx, req.BusinessId, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get wallets: %v", err)
	}

	var pbWallets []*pb.Wallet
	for _, w := range wallets {
		
		  addresses := w.Addresses()


		for _, addr := range addresses {
			pbWallets = append(pbWallets, &pb.Wallet{
				WalletId:  w.ID,
				Address:   addr.Address,
				Network:   string(addr.Network),
				Type:      pb.WalletType(pb.WalletType_value[string(w.Type)]),
				CreatedAt: w.CreatedAt.Format(time.RFC3339),
			})
		}
	}

	return &pb.BusinessWalletResponse{
		Wallets: pbWallets,
		Message: "wallets retrieved successfully",
	}, nil
}

func (s *WalletGRPCServer) WalletBalance(ctx context.Context, req *pb.WalletBalanceRequest) (*pb.WalletBalanceResponse, error) {
	if req.BusinessId == "" && req.CustomerId == "" {
		return nil, fmt.Errorf("either business_id or customer_id must be provided")
	}

	var customerID *string
	if req.CustomerId != "" {
		customerID = &req.CustomerId
	}

	var network db.Network
if req.Network != "" {
    switch strings.ToLower(req.Network) {
    case "evm":
        network = db.NetworkEvm
    case "solana":
        network = db.NetworkSolana
    case "tron":
        network = db.NetworkTron
    default:
        return nil, fmt.Errorf("invalid network: %s", req.Network)
    }
} else {
    return nil, fmt.Errorf("network must be provided")
}


	// Fetch wallet balances map
	balancesMap, err := s.service.WalletAddresses(ctx, req.BusinessId, customerID,network)
	if err != nil {
		return nil, fmt.Errorf("failed to get wallets: %v", err)
	}

	// Convert to proto map
	pbBalancesMap := make(map[string]*pb.ChainBalances) // ChainBalances is a repeated field wrapper

	for walletAddr, chainBalances := range balancesMap {
		var pbChainBalances []*pb.ChainBalanceResult
		for _, cb := range chainBalances {
			pbChainBalances = append(pbChainBalances, &pb.ChainBalanceResult{
				ChainName: cb.ChainName,
				Usdc:      cb.USDC,
				Usdt:      cb.USDT,
				Message:     "", // optionally map cb.Error if non-nil
			})
		}
		pbBalancesMap[walletAddr] = &pb.ChainBalances{Balances: pbChainBalances}
	}

	return &pb.WalletBalanceResponse{
		Balances: pbBalancesMap,
		Message:  "wallet balances retrieved successfully",
	}, nil
}


func (s *WalletGRPCServer) WalletChainBalance(ctx context.Context, req *pb.WalletBalanceRequest) (*pb.WalletChainBalanceResponse, error) {
	if req.BusinessId == "" && req.CustomerId == "" {
		return nil, fmt.Errorf("either business_id or customer_id must be provided")
	}

	var customerID *string
	if req.CustomerId != "" {
		customerID = &req.CustomerId
	}

	if req.Chain == "" {
		return nil, fmt.Errorf("chain must be provided")
	}
	chain := strings.ToLower(req.Chain)

	if req.Network == "" {
		return nil, fmt.Errorf("network must be provided")
	}

	var network db.Network
	switch strings.ToLower(req.Network) {
	case "evm":
		network = db.NetworkEvm
	case "solana":
		network = db.NetworkSolana
	case "tron":
		network = db.NetworkTron
	default:
		return nil, fmt.Errorf("invalid network: %s", req.Network)
	}

	// Fetch wallet balance for that chain
	balancesMap, err := s.service.WalletChainAddresses(ctx, chain, req.BusinessId, customerID, network)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch wallet balances: %v", err)
	}

	// Map[string]*pb.ChainBalanceResult
	pbBalancesMap := make(map[string]*pb.ChainBalanceResult)

	for walletAddr, cb := range balancesMap {
		pbBalancesMap[walletAddr] = &pb.ChainBalanceResult{
			ChainName: cb.ChainName,
			Usdc:      cb.USDC,
			Usdt:      cb.USDT,
			Message:     "",
		}
	}

	return &pb.WalletChainBalanceResponse{
		Balances: pbBalancesMap,
		
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
