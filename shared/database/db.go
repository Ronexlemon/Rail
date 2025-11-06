package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/ronexlemon/rail/shared/config"
	
)

var DB *sql.DB

func Connect() *sql.DB{
	cfg:= config.LoadDBConfig()
	db, err := sql.Open("pgx",cfg.URL)
	if err !=nil{
		log.Fatalf("Could not open Postgress connection %v",err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(),5 *time.Second)
	defer cancel()
	if err := db.PingContext(ctx);err !=nil{
		log.Fatalf("Failed to connect to Database %v",err)
	}
	log.Println("Connected to db successfully")
	DB =db
	return DB
}

func ConnectPrisma()