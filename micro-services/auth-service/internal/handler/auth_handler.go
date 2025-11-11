package handler

import (
	"encoding/json"
	"net/http"

	"context"
	"fmt"
	"time"

	"github.com/ronexlemon/rail/micro-services/auth-service/events"
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/service"
	pb "github.com/ronexlemon/rail/micro-services/auth-service/proto"
	
	businesspb "github.com/ronexlemon/rail/micro-services/business-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type BusinessHandler struct{
	service *service.BusinessService
}

func NewBusinessHandler(service *service.BusinessService)*BusinessHandler{
	return &BusinessHandler{service: service}
}

func (h *BusinessHandler) RegisterBusinessHandler(w http.ResponseWriter, r *http.Request){
	var req RegistrationReq
	

	if err:= json.NewDecoder(r.Body).Decode(&req); err !=nil{
		http.Error(w,"Invalid Request body",http.StatusBadRequest)
		return
	}
	user, err := h.service.RegisterBusiness(req.Email,req.Pass)
	if err !=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	//fire kafka event
	events.PublishEvent(user.ID,user)
	

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(user)

}





type AuthServer struct {
    pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	// 1. Dial the business-service using grpc.DialContext (non-deprecated)
	connCtx, cancelConn := context.WithTimeout(ctx, 5*time.Second)
	defer cancelConn()

	conn, err := grpc.DialContext(
		connCtx,
		"business-service:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // replaces WithInsecure
		grpc.WithBlock(), 
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to business-service: %v", err)
	}
	defer conn.Close()

	client := businesspb.NewBusinessServiceClient(conn)

	// 2. Call the business service
	resp, err := client.GetBusinessByKeys(ctx, &businesspb.GetBusinessByKeysRequest{
		ApiKey:    req.ApiKey,
		SecretKey: req.SecretKey,
	})
	if err != nil {
		return &pb.AuthenticateResponse{Valid: false}, nil
	}

	// 3. Map the response to a local Business model
	business := Business{
		BusinessID: resp.BusinessId,
		Status:     resp.Status,
		ApiKey:     req.ApiKey,
		SecretKey:  req.SecretKey,
	}

	
	return &pb.AuthenticateResponse{
		BusinessId: business.BusinessID,
		Status:     business.Status,
		Valid:      true,
	}, nil
}
