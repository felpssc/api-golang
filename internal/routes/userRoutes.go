package routes

import (
	createUserController "github.com/felpssc/api-golang/internal/modules/users/useCases/createUser"
	deleteUserController "github.com/felpssc/api-golang/internal/modules/users/useCases/deleteUser"
	getUserByIdController "github.com/felpssc/api-golang/internal/modules/users/useCases/getUserById"
	listUsersController "github.com/felpssc/api-golang/internal/modules/users/useCases/listUsers"
	updatedUserController "github.com/felpssc/api-golang/internal/modules/users/useCases/updateUser"
	"github.com/labstack/echo/v4"
)

func userRoutes(e *echo.Echo) {
	listUsersController.NewListUsersController(e)
	getUserByIdController.NewGetUserByIdController(e)
	createUserController.NewCreateUserController(e)
	updatedUserController.NewUpdateUserController(e)
	deleteUserController.NewDeleteUserController(e)
}
