package initializers

import (
	"minimalist-calories-api/errorHandling"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	errorHandling.FatalCheck("Error loading .env file", err)
}
