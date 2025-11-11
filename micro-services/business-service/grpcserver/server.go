package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ronexlemon/rail/micro-services/business-service/prisma/db"
	"github.com/yourusername/project/services/business-service/prisma/client"
	pb "github.com/yourusername/project/services/business-service/proto"
	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedBusinessServiceServer
}

func (s *server) GetBusinessByKeys(ctx context.Context, req *pb.GetBusinessByKeysRequest) (*pb.GetBusinessByKeysResponse, error) {
    

    business, err := db.Business.FindFirst(
        db.Business.ApiKey.Equals(req.ApiKey),
        db.Business.SecretKey.Equals(req.SecretKey),
    ).Exec(ctx)
    if err != nil {
        return nil, fmt.Errorf("invalid api key or secret")
    }

    return &pb.GetBusinessByKeysResponse{
        BusinessId: business.BusinessId,
        Status:     string(business.Status),
    }, nil
}

func GrpcServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterBusinessServiceServer(s, &server{})
    log.Println("Business service running on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
