package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupConnection() {
	dsn := "host=localhost user=postgres password=root dbname=go_crud port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Author{})
	// database.AutoMigrate(&Blog{})
	DB = database
}
