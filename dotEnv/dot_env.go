package dotEnv

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	// err := godotenv.Load("dotEnv/.env")
	err := godotenv.Load("C:/Users/ethmore/Projects/e-comm/auth-and-db-service/dotEnv/.env")

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}
