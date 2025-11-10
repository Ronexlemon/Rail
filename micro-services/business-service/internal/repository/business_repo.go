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


func (r *BusinessRepository) CreateBusiness(input BusinessInput)(*db.BusinessModel,error){

	business, err := r.Client.Business.CreateOne(
		db.Business.Name.Set(input.Name),
		db.Business.Email.Set(input.Email),
		//db.Business.BusinessID.Set(input.BusinessId),
		db.Business.APIKey.Set(input.ApiKey),
		db.Business.SecretKey.Set(input.SecretKey),

		
		
	).Exec(r.Context)

	if err != nil {
		return nil, err
	}
	fmt.Println("BUSINESS CREATED _________")
	return business, nil
}

func (r *BusinessRepository) CreateNewBusiness(input BusinessCreateInput)(*db.BusinessModel,error){

	business, err := r.Client.Business.CreateOne(
		db.Business.Name.Set(input.Name),
		db.Business.Email.Set(input.Email),
		//db.Business.BusinessID.Set(input.BusinessId),
		db.Business.APIKey.Set(input.ApiKey),
		db.Business.SecretKey.Set(input.SecretKey),

		
		
	).Exec(r.Context)

	if err != nil {
		return nil, err
	}
	fmt.Println("BUSINESS CREATED _________")
	return business, nil
}