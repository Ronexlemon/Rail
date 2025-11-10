package service

import (
	"fmt"

	"github.com/ronexlemon/rail/micro-services/business-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/business-service/prisma/db"
	"github.com/ronexlemon/rail/micro-services/business-service/utils"
)


type BusinessService struct{
	repo *repository.BusinessRepository
}

func NewBusinesssService(repo *repository.BusinessRepository)*BusinessService{
	return &BusinessService{
		repo: repo,
	}
}

func (s *BusinessService) CreateBusinessService(name,email,businessId string)(*db.BusinessModel,error){
	if name == "" || email == "" || businessId == ""{

		return nil,fmt.Errorf("email and name, businessId are required")
	}
	result,_ :=utils.GenerateAPIKeys(name)
	 input := repository.BusinessInput{
		Email: email,
		BusinessId: businessId,
		ApiKey: result.PublicKey,
		SecretKey: result.SecretKey,
	}
	
	return s.repo.CreateBusiness(input)
}