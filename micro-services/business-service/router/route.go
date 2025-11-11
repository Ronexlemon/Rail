package router

import 

( "github.com/ronexlemon/rail/micro-services/business-service/internal/handler"

  "github.com/gorilla/mux"

)



func NewRouter(h *handler.BusinessServiceHandler)*mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/v0/register-business",h.RegisterBusinessHandler).Methods("POST")
	r.HandleFunc("/v0/customer/{customerId}/wallet",h.RegisterBusinessHandler).Methods("POST")

	return r
}