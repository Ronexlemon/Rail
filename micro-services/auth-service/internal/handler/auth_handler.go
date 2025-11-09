package handler

import (
	"encoding/json"
	"net/http"

	
	"github.com/ronexlemon/rail/micro-services/auth-service/events"
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/service"
	
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