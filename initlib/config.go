package initlib

import "os"

type Config struct {
	DBHost     string
	DBUserName string
	DBUserPwd  string
	DBName     string
	DBPort     string
	ClientOrigin string
}


func LoadConfig() (config Config) {
	return Config{
		DBHost: os.Getenv("MYSQL_HOST"),
		DBUserName: os.Getenv("MYSQL_USER"),
		DBUserPwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		DBPort: os.Getenv("MYSQL_PORT"),
		ClientOrigin: os.Getenv("CLIENT_ORIGIN"),
	}
}