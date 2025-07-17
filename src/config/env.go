package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type env struct {
	Port          string
	DbPostgresUrl string
}

var Env env

func Init() {
	if err := godotenv.Load(); err == nil {
		log.Printf("==> EnvFile found!!!")
	}

	Env.Port = os.Getenv("PORT")
	Env.DbPostgresUrl = os.Getenv("DB_POSTGRES_URL")
}
