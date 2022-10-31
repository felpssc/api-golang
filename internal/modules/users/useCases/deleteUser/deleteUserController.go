package usecase

import (
	"strconv"

	IusersRepository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository"
	repository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository/implementations"
	utils "github.com/felpssc/api-golang/pkg/utils/repositoryManager"
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

	usersRepository := utils.GetRepository[IusersRepository.UsersRepository](
		repository.NewPostgresUsersRepository(),
		repository.NewLocalUsersRepository(),
	)

	err = NewDeleteUserUseCase(usersRepository).execute(id)

	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, "User deleted")
}
