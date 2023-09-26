package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	Port          string
	JWTSecret     string
	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        string
	MailgunDomain string
	MailgunKey    string
	MailSender    string
}

var Env *EnvVars

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Env = &EnvVars{
		Port:          os.Getenv("PORT"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		DBHost:        os.Getenv("DB_HOST"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBPort:        os.Getenv("DB_PORT"),
		MailgunDomain: os.Getenv("MAILGUN_DOMAIN"),
		MailgunKey:    os.Getenv("MAILGUN_KEY"),
		MailSender:    os.Getenv("MAIL_SENDER"),
	}
}
