package handler



type RegistrationReq struct{
	Email string `json:"email"`
	
	Pass  string `json:"pass"`
}

type Business struct{
		BusinessID string
		Status     string
		ApiKey     string
		SecretKey  string
	}