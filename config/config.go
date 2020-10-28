package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT                  = 0
	BIND_ADDRESS          = ""
	DB_DRIVER             = ""
	DB_USER               = ""
	DB_PASS               = ""
	DB_NAME               = ""
	DB_PORT               = 0
	SSL_MODE              = ""
	TIMEZONE              = ""
	DB_URL                = ""
	COGNITO_CLIENT_ID     = ""
	COGNITO_CLIENT_SECRET = ""
	COGNITO_USER_POOL_ID  = ""
	AWS_REGION            = "eu-central-1"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Http Server

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 8000
	}

	if os.Getenv("BIND_ADDRESS") == "" {
		BIND_ADDRESS = "0.0.0.0"
	} else {
		BIND_ADDRESS = os.Getenv("BIND_ADDRESS")
	}

	// Database
	DB_DRIVER = os.Getenv("DB_DRIVER")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")

	DB_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		DB_PORT = 5432
	}

	SSL_MODE = os.Getenv("SSL_MODE")
	TIMEZONE = os.Getenv("TIMEZONE")

	// DB_URL form the connection string for postgresql database
	// "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	DB_URL = fmt.Sprintf("host=0.0.0.0 user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("SSL_MODE"), os.Getenv("TIMEZONE"))

	COGNITO_CLIENT_ID = os.Getenv("COGNITO_CLIENT_ID")
	COGNITO_CLIENT_SECRET = os.Getenv("COGNITO_CLIENT_SECRET")
	COGNITO_USER_POOL_ID = os.Getenv("COGNITO_USER_POOL_ID")

	AWS_REGION = os.Getenv("AWS_REGION")

}
