package grserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ronexlemon/rail/micro-services/business-service/database"
	"github.com/ronexlemon/rail/micro-services/business-service/prisma/db"
	pb "github.com/ronexlemon/rail/micro-services/business-service/proto"
	"google.golang.org/grpc"
)

type BusinessRepository struct {
	Client  *db.PrismaClient
	Context context.Context
}

func NewBusinessRepository() *BusinessRepository {
	return &BusinessRepository{
		Client:  database.PrismaDBClient.Client,
		Context: database.PrismaDBClient.Context,
	}
}

type Server struct {
	pb.UnimplementedBusinessServiceServer
	repo *BusinessRepository
}

func NewServer(repo *BusinessRepository) *Server {
	return &Server{repo: repo}
}

func (s *Server) GetBusinessByKeys(ctx context.Context, req *pb.GetBusinessByKeysRequest) (*pb.GetBusinessByKeysResponse, error) {
	// Example using Prisma Go client (adjust for your schema)
	business, err := s.repo.Client.Business.FindFirst(
		db.Business.APIKey.Equals(req.ApiKey),
		db.Business.SecretKey.Equals(req.SecretKey),
	).Exec(ctx)
	if err != nil || business == nil {
		return nil, fmt.Errorf("invalid api key or secret")
	}

	return &pb.GetBusinessByKeysResponse{
		BusinessId: business.ID,
		Status:     string(business.Status),
	}, nil
}

func GrpServer() {
	repo := NewBusinessRepository()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBusinessServiceServer(s, NewServer(repo))
	log.Println("Business service running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
