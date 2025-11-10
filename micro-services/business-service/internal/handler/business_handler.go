package handler

import (
	

	"github.com/ronexlemon/rail/micro-services/business-service/internal/service"
)


type BusinessServiceHandler struct{
	service *service.BusinessService
}
func NewBusinessService(service *service.BusinessService)*BusinessServiceHandler{
	return  &BusinessServiceHandler{
		service: service,

	}
}



// func (h *BusinessServiceHandler) RegisterBusinessHandler(w http.ResponseWriter, r *http.Request){
// 	var req RegistrationReq
	

// 	if err:= json.NewDecoder(r.Body).Decode(&req); err !=nil{
// 		http.Error(w,"Invalid Request body",http.StatusBadRequest)
// 		return
// 	}
// 	user, err := h.service.RegisterBusiness(req.Email,req.Pass)
// 	if err !=nil{
// 		http.Error(w,err.Error(),http.StatusInternalServerError)
// 		return
// 	}

// 	//fire kafka event
// 	events.PublishEvent(user.ID,user)
	

// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(user)

// }