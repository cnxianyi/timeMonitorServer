package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitEnv() {

	envFile := os.Getenv("GIN_MODE")

	if envFile != "" {
		err := godotenv.Load(".env." + envFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}
}
