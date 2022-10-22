package usecase

import (
	"strconv"

	repository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository/implementations"
	"github.com/labstack/echo/v4"
)

func NewGetUserByIdController(e *echo.Echo) {
	e.GET("/users/:id", getUserByIdUseCase)
}

func getUserByIdUseCase(c echo.Context) error {

	idString := c.Param("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		return c.JSON(400, "Invalid id")
	}

	usersRepository := repository.NewLocalUsersRepository()

	user, err := NewGetUserByIdUseCase(usersRepository).execute(id)

	if err != nil {
		return c.JSON(200, err.Error())
	}

	return c.JSON(200, user)
}
