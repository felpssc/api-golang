package updateUserUseCase

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewUpdateUserController(e *echo.Echo) {
	e.PUT("/users/:id", updateUserUseCase)
}

func updateUserUseCase(c echo.Context) error {

	idString := c.Param("id")

	var nameJson map[string]interface{} = map[string]interface{}{}

	if err := c.Bind(&nameJson); err != nil {
		return c.JSON(400, "Invalid json")
	}

	name := nameJson["name"].(string)

	id, err := strconv.Atoi(idString)

	if err != nil {
		return c.JSON(400, "Invalid id")
	}

	user, err := UpdateUserUseCase(id, name)

	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, user)
}
