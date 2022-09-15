package config

import (
	"gift-app/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var dbUser = os.Getenv("DBUSER")
var dbPass = os.Getenv("DBPASS")
var dbHost = os.Getenv("DBHOST")
var dbPort = os.Getenv("DBPORT")

var Database *gorm.DB

func Connect() error {
	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" {
		panic("Missing environment variables for database connection.")
		return nil
	}

	var DatabaseUri = dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	Database, err = gorm.Open(mysql.Open(DatabaseUri), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Gift{})

	return nil
}
