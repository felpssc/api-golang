package usecase

import (
	IusersRepository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository"
	repository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository/implementations"
	utils "github.com/felpssc/api-golang/pkg/utils/repositoryManager"
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
	password := json["password"].(string)

	usersRepository := utils.GetRepository[IusersRepository.UsersRepository](
		repository.NewPostgresUsersRepository(),
		repository.NewLocalUsersRepository(),
	)

	createdUser, err := NewCreateUserUseCase(usersRepository).execute(name, email, password)

	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, createdUser)
}
