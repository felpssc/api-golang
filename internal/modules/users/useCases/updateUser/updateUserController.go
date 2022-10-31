package usecase

import (
	"strconv"

	IusersRepository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository"
	repository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository/implementations"
	utils "github.com/felpssc/api-golang/pkg/utils/repositoryManager"
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

	usersRepository := utils.GetRepository[IusersRepository.UsersRepository](
		repository.NewPostgresUsersRepository(),
		repository.NewLocalUsersRepository(),
	)

	user, err := NewUpdateUserUseCase(usersRepository).execute(id, name)

	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, user)
}
