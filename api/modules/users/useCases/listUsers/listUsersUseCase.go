package listUsersUseCase

import (
	user "github.com/felpssc/api-golang/api/modules/users/entities/user"
	usersRepository "github.com/felpssc/api-golang/api/modules/users/repositories/usersRepository"
)

func ListUsersUseCase() []user.User {
	users := usersRepository.GetUsers()
	return users
}
