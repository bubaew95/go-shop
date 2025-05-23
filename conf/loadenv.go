package conf

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvOptional(envPath string) error {
	if envPath == "" {
		envPath = ".env"
	}

	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		return err
	}

	if err := godotenv.Load(envPath); err != nil {
		return err
	}

	return nil
}
