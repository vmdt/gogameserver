package system

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	} else {
		log.Println("No .env file found, using default environment variables")
	}
}

func Getenv(key string, args ...string) string {
	def := ""
	if len(args) > 0 {
		def = args[0]
	}
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func init() {
	loadEnv()
}
