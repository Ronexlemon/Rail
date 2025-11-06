package config

import "os"

type DBConfig struct{
	URL string
}

func LoadDBConfig() *DBConfig{
	url := GetEnv("DATABASE_URL","local")

	return &DBConfig{URL: url}
}


func GetEnv(key string,fallback string)string{
	if val := os.Getenv(key); val !=""{
		return val
	}
	return fallback
}
