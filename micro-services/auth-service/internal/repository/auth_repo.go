package repository

import (
	"context"
	"github.com/ronexlemon/rail/micro-services/auth-service/prisma/db"
	"github.com/ronexlemon/rail/micro-services/auth-service/database"
)

type BusinessRepository struct {
	client  *db.PrismaClient
	context context.Context
}

// Constructor
func NewBusinessRepository() *BusinessRepository {

	
	return &BusinessRepository{
		client:  database.PrismaDBClient.Client,
		context: database.PrismaDBClient.Context,
	}
}

// CreateBusiness creates a new business user
func (r *BusinessRepository) CreateBusiness(email, companyReg ,name string) (*db.UserModel, error) {
	user, err := r.client.User.CreateOne(
		db.User.Name.Set(name),         
		db.User.Email.Set(email),        
		db.User.CompanyReg.Set(companyReg), 
		db.User.Type.Set(db.RoleBusiness), 
	).Exec(r.context)

	if err != nil {
		return nil, err
	}
	return user, nil
}
