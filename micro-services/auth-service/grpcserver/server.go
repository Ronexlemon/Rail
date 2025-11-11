package grpcserver

import (
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/auth-service/proto"
)





type AuthGRPCServer struct{
	proto.UnimplementedAuthServiceServer
	repo *repository.BusinessRepository
}

func NewServer(repo *BusinessRepository) *AuthGRPCServer {
	return &AuthGRPCServer{repo: repo}
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

func GrpcServer() {
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
