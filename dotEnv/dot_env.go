package dotEnv

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	err := godotenv.Load("dotEnv/.env")

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}
