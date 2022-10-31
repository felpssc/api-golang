package config

import (
	"fmt"

	entities "github.com/felpssc/api-golang/internal/modules/users/entities/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")

	db = database
}

func GetDB() *gorm.DB {
	return db
}

func init() {
	Connect()
	db = GetDB()
	db.AutoMigrate(&entities.User{})

	fmt.Println("Database migrated")
}
