package handler

import (
	"encoding/json"
	"net/http"

	
	"github.com/ronexlemon/rail/micro-services/auth-service/events"
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/service"
	 "context"
    "fmt"
    pb "github.com/ronexlemon/rail/micro-services/auth-service/proto"
    businesspb "github.com/ronexlemon/rail/micro-services/business-service/internal/handler"
    "google.golang.org/grpc"
    "time"
	
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
    conn, err := grpc.Dial("business-service:50051", grpc.WithInsecure())
    if err != nil {
        return nil, fmt.Errorf("failed to connect to business-service: %v", err)
    }
    defer conn.Close()

    client := businesspb.NewBusinessServiceClient(conn)
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    resp, err := client.GetBusinessByKeys(ctx, &businesspb.GetBusinessByKeysRequest{
        ApiKey:    req.ApiKey,
        SecretKey: req.SecretKey,
    })
    if err != nil {
        return &pb.AuthenticateResponse{Valid: false}, nil
    }

    return &pb.AuthenticateResponse{
        BusinessId: resp.BusinessId,
        Status:     resp.Status,
        Valid:      true,
    }, nil
}
