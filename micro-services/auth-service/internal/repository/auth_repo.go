package repository

import (
	"context"
	"fmt"

	"github.com/ronexlemon/rail/micro-services/auth-service/database"
	"github.com/ronexlemon/rail/micro-services/auth-service/prisma/db"
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
func (r *BusinessRepository) CreateBusiness(email ,name,passwordHash ,apiKey,secretKey string) (*db.UserModel, error) {

	user, err := r.client.User.CreateOne(
		db.User.Email.Set(email),
		db.User.PasswordHash.Set(passwordHash),         
		db.User.Role.Set("BUSINESS"),        
		db.User.APIKey.Set(apiKey),
		db.User.SecretKey.Set(secretKey),		

	).Exec(r.context)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *BusinessRepository) ValidateAPIKeys(apiKey, secretKey string) (*db.UserModel, error) {
	
	user, err := r.client.User.FindFirst(
		db.User.APIKey.Equals(apiKey),
		db.User.SecretKey.Equals(secretKey),
	).Exec(r.context)

	if err != nil {
		
		return nil, fmt.Errorf("invalid apiKey or secretKey: %v", err)
	}

	if user == nil {
		return nil, fmt.Errorf("invalid apiKey or secretKey")
	}

	return user, nil
}
