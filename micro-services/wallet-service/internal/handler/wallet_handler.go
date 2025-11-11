package handler

import (
	"net/http"

	"github.com/ronexlemon/rail/micro-services/wallet-service/internal/service"
)


type WalletHandler struct{
	service *service.WalletService
}

func NewWalletHandler(service *service.WalletService)*WalletHandler{
	return &WalletHandler{
		service: service,
	}
}

// func (h *BusinessServiceHandler) CustomerWalletHandler(w http.ResponseWriter, r *http.Request){
// 	var req CustomerWalletInput
	

// 	if err:= json.NewDecoder(r.Body).Decode(&req); err !=nil{
// 		http.Error(w,"Invalid Request body",http.StatusBadRequest)
// 		return
// 	}
// 	user, err := h.service.CreateNewBusinessService(req.Name,req.Email)
// 	if err !=nil{
// 		http.Error(w,err.Error(),http.StatusInternalServerError)
// 		return
// 	}

// 	//fire kafka event
// 	events.PublishEvent(user.ID,user)
	

// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(user)

// }


func (h *WalletHandler) CustomerWalletCreationHandler(w http.ResponseWriter, r *http.Request){
 //var req CustomerWallet //has Customer_id in struct


}