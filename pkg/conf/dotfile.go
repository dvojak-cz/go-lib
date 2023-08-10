package conf

import (
	"os"
	"github.com/joho/godotenv"
)
func getEnvFilePath() string {
	var envPath string
	if envPath = os.Getenv("ENVIRONMENT_FILE"); envPath == "" {
		envPath = ".env"
	}
	return envPath
}

func init() {
	envPath := getEnvFilePath()
	godotenv.Load(envPath)
}