package updateUser

import (
	"errors"

	user "github.com/felpssc/api-golang/api/modules/users/entities/user"
	usersRepository "github.com/felpssc/api-golang/api/modules/users/repositories/usersRepository"
)

func UpdateUserUseCase(id int, name string) (user.User, error) {
	user, err := usersRepository.GetUserByID(id)

	if err != nil {
		return user, errors.New("user not found")
	}

	updatedUser, err := usersRepository.UpdateUser(id, name)

	if err != nil {
		return user, err
	}

	return updatedUser, nil
}
