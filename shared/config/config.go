package config

import "os"

type DBConfig struct{
	URL string
}

func LoadDBConfig() *DBConfig{
	url := getEnv("DATABASE_URL",nil)

	return &DBConfig{URL: url}
}


func getEnv(key string,fallback *string)string{
	if val := os.Getenv(key); val !=""{
		return val
	}
	return *fallback
}
