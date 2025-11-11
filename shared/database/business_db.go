package database

import (
    "context"
    "fmt"
    "log"

   
    "github.com/ronexlemon/rail/micro-services/business-service/internal/handler"
)

var PrismaDBClient *db.PrismaClient
var Ctx context.Context

func ConnectBusinessDB() (*db.PrismaClient, error) {
    client := db.NewClient()
    ctx := context.Background()

    if err := client.Connect(); err != nil {
        return nil, fmt.Errorf("failed to connect to business DB: %w", err)
    }

    PrismaDBClient = client
    Ctx = ctx

    log.Println("Business DB connected successfully")
    return client, nil
}
