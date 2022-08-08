package deleteUserUseCase

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewDeleteUserController(e *echo.Echo) {
	e.DELETE("/users/:id", deleteUserUseCase)
}

func deleteUserUseCase(c echo.Context) error {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		return c.JSON(400, "invalid id")
	}

	err = DeleteUserUseCase(id)

	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, "User deleted")
}
