package repository

import (
	"context"
	"fmt"

	"github.com/ronexlemon/rail/micro-services/business-service/database"
	"github.com/ronexlemon/rail/micro-services/business-service/prisma/db"
)


type BusinessRepository struct{
	Client *db.PrismaClient
	Context context.Context
}

func NewBusinessRepository()(*BusinessRepository){
	return &BusinessRepository{
		Client: database.PrismaDBClient.Client,
		Context: database.PrismaDBClient.Context,
	}
}

type BusinessInput struct {
	Name       string
	Email      string
	BusinessId string // from Auth service
	ApiKey    string
	SecretKey  string
	
}
type BusinessCreateInput struct{
	Name  string
    Email string 
    ApiKey  string 
    SecretKey  string
    BusinessId string  
}


type BusinessData struct{
	ID  string `json:"auth_user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func (r *BusinessRepository) CreateBusiness(input BusinessData)(*db.BusinessModel,error){

	business, err := r.Client.Business.CreateOne(
		db.Business.AuthUserID.Set(input.ID),
		db.Business.Name.Set(input.Name),
		db.Business.Email.Set(input.Email),
	).Exec(r.Context)

	if err != nil {
		return nil, err
	}
	fmt.Println("BUSINESS CREATED _________")
	return business, nil
}

func (r *BusinessRepository) CreateNewBusiness(input BusinessData)(*db.BusinessModel,error){

	business, err := r.Client.Business.CreateOne(
		db.Business.AuthUserID.Set(input.ID),
		db.Business.Name.Set(input.Name),
		db.Business.Email.Set(input.Email),
	).Exec(r.Context)
	if err != nil {
		return nil, err
	}
	fmt.Println("BUSINESS CREATED _________")
	return business, nil
}