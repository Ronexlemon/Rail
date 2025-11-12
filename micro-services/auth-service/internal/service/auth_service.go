package service

import (
	"fmt"

	
	"github.com/ronexlemon/rail/micro-services/auth-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/auth-service/prisma/db"
	"github.com/ronexlemon/rail/micro-services/auth-service/utils"
)


type BusinessService struct{
	repo *repository.BusinessRepository
}

func NewBusinessService(repo *repository.BusinessRepository)*BusinessService{
	return &BusinessService{repo: repo}
}

func(s *BusinessService) RegisterBusiness(email,name,pass string)(*db.UserModel,error){
	//todo hash password
	
	if email == "" || pass == ""{
		return nil,fmt.Errorf("email and pass are required")
	}
	result,err := utils.GenerateAPIKeys()
 
	if err !=nil{
		return nil,fmt.Errorf("failed to create auth keys")
	}
	apiKey := result.PublicKey
	secretKey := result.SecretKey
	return s.repo.CreateBusiness(email,name,pass,apiKey,secretKey)
}