package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DbName                 string
	DbUser                 string
	DbPassword             string
	DbHost                 string
	DbPort                 string
	ServerPort             string
	EmailForEmailToSend    string
	PasswordForEmailToSend string
	ServerNgrokHost        string
	//SERVER_NGROK_HOST
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) InitENV() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg.DbName = os.Getenv("DB_NAME")
	cfg.DbUser = os.Getenv("DB_USER")
	cfg.DbPassword = os.Getenv("DB_PASSWORD")
	cfg.ServerPort = os.Getenv("SERVER_PORT")
	cfg.EmailForEmailToSend = os.Getenv("EMAIL_FOR_EMAIL_TO_SEND")
	cfg.PasswordForEmailToSend = os.Getenv("PASSWORD_FOR_EMAIL_TO_SEND")
	cfg.DbHost = os.Getenv("DB_HOST")
	cfg.DbPort = os.Getenv("DB_PORT")
	cfg.ServerNgrokHost = os.Getenv("SERVER_NGROK_HOST")
}
