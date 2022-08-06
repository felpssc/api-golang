package listUsersUseCase

import (
	"github.com/labstack/echo/v4"
)

func NewListUsersController(e *echo.Echo) {
	e.GET("/users", listUsersUseCase)
}

func listUsersUseCase(c echo.Context) error {
	return c.JSON(200, ListUsersUseCase())
}
