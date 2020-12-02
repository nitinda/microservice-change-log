package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	API_PORT         = 0
	API_BIND_ADDRESS = ""
	API_DB_HOST      = ""
	API_DB_DRIVER    = ""
	API_DB_USER      = ""
	API_DB_PASS      = ""
	API_DB_NAME      = ""
	API_DB_PORT      = 5432
	API_SSL_MODE     = "disabled"
	API_TIMEZONE     = ""
	API_DB_URL       = ""
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Http Server

	// var err error

	API_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		API_PORT = 8000
	}

	if os.Getenv("API_BIND_ADDRESS") == "" {
		API_BIND_ADDRESS = "0.0.0.0"
	} else {
		API_BIND_ADDRESS = os.Getenv("API_BIND_ADDRESS")
	}

	// Database
	API_DB_DRIVER = os.Getenv("API_DB_DRIVER")
	API_DB_USER = os.Getenv("API_DB_USER")
	API_DB_PASS = os.Getenv("API_DB_PASS")
	API_DB_NAME = os.Getenv("API_DB_NAME")

	API_DB_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		API_DB_PORT = 5432
	}

	API_SSL_MODE = os.Getenv("API_SSL_MODE")
	API_TIMEZONE = os.Getenv("API_TIMEZONE")

	// DB_URL form the connection string for postgresql database
	// "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	API_DB_URL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("API_DB_HOST"), os.Getenv("API_DB_USER"), os.Getenv("API_DB_PASS"),
		os.Getenv("API_DB_NAME"), os.Getenv("API_DB_PORT"), os.Getenv("API_SSL_MODE"),
		os.Getenv("API_TIMEZONE"))
}
