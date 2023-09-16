package utils

import "os"

type Env struct {
	Port          string
	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        string
	MailgunDomain string
	MailgunKey    string
}

func GetEnv() Env {
	env := Env{
		Port:          os.Getenv("PORT"),
		DBHost:        os.Getenv("DB_HOST"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBPort:        os.Getenv("DB_PORT"),
		MailgunDomain: os.Getenv("MAILGUN_DOMAIN"),
		MailgunKey:    os.Getenv("MAILGUN_KEY"),
	}
	return env
}
