package initlib

import (
	"fmt"
	"github.com/arencloud/space-demo/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase(config *Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUserName, config.DBUserPwd, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database! \n", err.Error())
		os.Exit(1)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	db.AutoMigrate(&models.Note{})

	log.Println("Connected Successfully to the Database!!!")
}
