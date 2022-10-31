package main

import (
	"fmt"

	router "github.com/felpssc/api-golang/internal/routes"
	godotenv "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	e := echo.New()

	router.Router(e)

	e.Logger.Fatal(e.Start(":3000"))
}
