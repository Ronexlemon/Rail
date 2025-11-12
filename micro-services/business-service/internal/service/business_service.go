package service

import (
	"fmt"
	"log"

	"github.com/ronexlemon/rail/micro-services/business-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/business-service/prisma/db"
	
)


type BusinessService struct{
	repo *repository.BusinessRepository
}

func NewBusinesssService(repo *repository.BusinessRepository)*BusinessService{
	return &BusinessService{
		repo: repo,
	}
}

func (s *BusinessService) CreateBusinessService(name,email,businessId,authUserId string)(*db.BusinessModel,error){
	log.Printf("Value create Key=%s Value=%s  =%s", string(name), string(email),string(businessId))
	if name == "" || email == "" || businessId == ""{

		return nil,fmt.Errorf("email and name, businessId are required")
	}
	//result,_ :=utils.GenerateAPIKeys(name)
	 input := repository.BusinessData{
		Email: email,
		Name: name,
       ID: authUserId,
		
	}
	
	return s.repo.CreateBusiness(input)
}

func (s *BusinessService) CreateNewBusinessService(name,email,authUserId string)(*db.BusinessModel,error){
	log.Printf("Value create Key=%s Value=%s  =%s", string(name), string(email))
	if name == "" || email == "" {

		return nil,fmt.Errorf("email and name are required")
	}
	//result,_ :=utils.GenerateAPIKeys(name)
	 input := repository.BusinessData{
		Email: email,
		Name: name,
       ID: authUserId,
		
	}
	
	return s.repo.CreateBusiness(input)
}