package database

import (
	"context"
	
	"github.com/ronexlemon/rail/micro-services/business-service/prisma/db"
)

type PrismaDB struct {
	Client *db.PrismaClient
	Context context.Context
}
var PrismaDBClient = &PrismaDB{}

func ConnectDB()(*PrismaDB,error){
	
	
	client := db.NewClient()
	if err := client.Prisma.Connect(); err !=nil{
		return nil,err
	}
	PrismaDBClient.Client = client
	PrismaDBClient.Context= context.Background()
	return PrismaDBClient,nil
}