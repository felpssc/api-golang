package main

import (
	router "github.com/felpssc/api-golang/api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	router.Router(e)

	e.Logger.Fatal(e.Start(":3000"))
}
