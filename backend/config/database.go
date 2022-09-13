package config

import (
	"gift-app/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DatabaseUri = "root:passwd@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
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
