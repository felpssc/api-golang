package createUserUseCase

import (
	"github.com/labstack/echo/v4"
)

func NewCreateUserController(e *echo.Echo) {
	e.POST("/users", createUserUseCase)
}

func createUserUseCase(c echo.Context) error {

	var json map[string]interface{} = map[string]interface{}{}

	if err := c.Bind(&json); err != nil {
		return c.JSON(400, "Invalid json")
	}

	name := json["name"].(string)
	email := json["email"].(string)

	createdUser, err := CreateUserUseCase(name, email)

	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, createdUser)
}