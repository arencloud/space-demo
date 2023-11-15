package initlib

import "os"

type Config struct {
	DBHost       string
	DBUserName   string
	DBUserPwd    string
	DBName       string
	DBPort       string
	ClientOrigin string
}

func LoadConfig() (config Config) {
	return Config{
		DBHost:       os.Getenv("SPACE_DEMO_DB_HOST"),
		DBUserName:   os.Getenv("SPACE_DEMO_DB_USER"),
		DBUserPwd:    os.Getenv("SPACE_DEMO_DB_PASSWORD"),
		DBName:       os.Getenv("SPACE_DEMO_DB_DATABASE"),
		DBPort:       os.Getenv("SPACE_DEMO_DB_PORT"),
		ClientOrigin: os.Getenv("CLIENT_ORIGIN"),
	}
}
