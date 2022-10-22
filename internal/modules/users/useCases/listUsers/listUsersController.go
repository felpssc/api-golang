package usecase

import (
	repository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository/implementations"
	"github.com/labstack/echo/v4"
)

func NewListUsersController(e *echo.Echo) {
	e.GET("/users", listUsersUseCase)
}

func listUsersUseCase(c echo.Context) error {
	usersRepository := repository.NewLocalUsersRepository()

	users := NewListUsersUseCase(usersRepository).execute()

	return c.JSON(200, users)
}
