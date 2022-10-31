package usecase

import (
	IusersRepository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository"
	repository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository/implementations"
	utils "github.com/felpssc/api-golang/pkg/utils/repositoryManager"
	"github.com/labstack/echo/v4"
)

func NewListUsersController(e *echo.Echo) {
	e.GET("/users", listUsersUseCase)
}

func listUsersUseCase(c echo.Context) error {
	usersRepository := utils.GetRepository[IusersRepository.UsersRepository](
		repository.NewPostgresUsersRepository(),
		repository.NewLocalUsersRepository(),
	)

	users := NewListUsersUseCase(usersRepository).execute()

	return c.JSON(200, users)
}
