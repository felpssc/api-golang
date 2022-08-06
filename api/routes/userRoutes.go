package routes

import (
	createUserController "github.com/felpssc/api-golang/api/modules/users/useCases/createUser"
	getUserByIdController "github.com/felpssc/api-golang/api/modules/users/useCases/getUserById"
	listUsersController "github.com/felpssc/api-golang/api/modules/users/useCases/listUsers"
	"github.com/labstack/echo/v4"
)

func userRoutes(e *echo.Echo) {
	listUsersController.NewListUsersController(e)
	getUserByIdController.NewGetUserByIdController(e)
	createUserController.NewCreateUserController(e)
}
