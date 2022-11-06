package config

import (
	"fmt"
	"os"

	entities "github.com/felpssc/api-golang/internal/modules/users/entities/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {

	err := godotenv.Load()

	if err != nil {
		panic("Failed to load .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
	)

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
