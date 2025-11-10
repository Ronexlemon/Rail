package handler



type RegistrationReq struct{
	Email string `json:"email"`
	
	Pass  string `json:"pass"`
}

