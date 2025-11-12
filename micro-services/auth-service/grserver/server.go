package grserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ronexlemon/rail/micro-services/auth-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/service"
	pb "github.com/ronexlemon/rail/micro-services/auth-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)




type AuthGRPCServer struct{
	pb.UnimplementedAuthServiceServer
	repo *repository.BusinessRepository
	service *service.BusinessService
}

func NewServer(repo *repository.BusinessRepository,service *service.BusinessService) *AuthGRPCServer {
	return &AuthGRPCServer{repo: repo,service: service}
}

func (s *AuthGRPCServer) GetBusinessByKeys(ctx context.Context, req *pb.GetBusinessByKeysRequest) (*pb.GetBusinessByKeysResponse, error) {
	
	business, err := s.repo.ValidateAPIKeys(req.ApiKey,req.SecretKey)
	if err != nil || business == nil {
		return nil, fmt.Errorf("invalid api key or secret")
	}

	return &pb.GetBusinessByKeysResponse{
		BusinessId: business.ID,
		Status:     string(business.Status),
	}, nil
}

func (s *AuthGRPCServer) RegisterBusiness(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	
	business, err := s.service.RegisterBusiness(req.Email, req.Name, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create business: %v", err)
	}

	if business == nil {
		return nil, status.Error(codes.Internal, "failed to create business, please try again")
	}

	return &pb.RegisterUserResponse{
		UserId:    business.ID,
		ApiKey:    business.APIKey,
		SecretKey: business.SecretKey,
		Status:    string(business.Status),
		Message:   "Business registered successfully",
	}, nil
}

func ServerGrpc() {
	repo := repository.NewBusinessRepository()
	service := service.NewBusinessService(repo)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, NewServer(repo,service))
	log.Println("Business service running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
