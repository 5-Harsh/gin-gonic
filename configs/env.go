package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	return os.Getenv("MONGO_URI")
}

func LoadEnvVariable(varName string) string {
	err := godotenv.Load()

	if err != nil {

		log.Fatal("Error Loading .env file")

	}
	return os.Getenv(varName)
}
