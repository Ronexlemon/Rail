package handler

type RegistrationReq struct{
	Email string `json:"email"`
	CompanyReg string `json:"company_reg"`
	name  string `json:"name"`
}