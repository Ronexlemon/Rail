package handler

type BusinessCreateInput struct{
	Name  string `json:"name"`
    Email string  `json:"email"`
}

type CustomerWalletInput struct{
	CustomerID string `json:"customer_id"` 
}