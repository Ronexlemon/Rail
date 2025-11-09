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
func (r *BusinessRepository) CreateBusiness(email ,name string) (*db.UserModel, error) {
// 	passwordHash  String
//   role  Role 
//   email      String   @unique
	user, err := r.client.User.CreateOne(
		db.User.PasswordHash.Set(name),         
		db.User.Role.Set("BUSINESS"),        
		db.User.Email.Set(email),
	).Exec(r.context)

	if err != nil {
		return nil, err
	}
	return user, nil
}
