package configs



import "os"






func GetEnv(key string,fallback string)string{
	if val := os.Getenv(key); val !=""{
		return val
	}
	return fallback
}
