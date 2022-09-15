package config

import (
	"gift-app/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DatabaseUrl = os.Getenv("DB_URL")

var Database *gorm.DB

func Connect() error {
	if DatabaseUrl == "" {
		panic("Missing environment variable for database connection.")
		return nil
	}

	var err error

	Database, err = gorm.Open(mysql.Open(DatabaseUrl), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Gift{})

	return nil
}
