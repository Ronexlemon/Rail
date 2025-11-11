package service

import (
	"fmt"

	"github.com/ronexlemon/rail/micro-services/auth-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/auth-service/prisma/db"
)


type BusinessService struct{
	repo *repository.BusinessRepository
}

func NewBusinessService(repo *repository.BusinessRepository)*BusinessService{
	return &BusinessService{repo: repo}
}

func(s *BusinessService) RegisterBusiness(email,pass string)(*db.UserModel,error){

	if email == "" || pass == ""{
		return nil,fmt.Errorf("email and pass are required")
	}
	return s.repo.CreateBusiness(email,"fff","fsf","sfs")
}