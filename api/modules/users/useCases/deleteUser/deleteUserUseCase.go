package deleteUserUseCase

import (
	"errors"

	usersRepository "github.com/felpssc/api-golang/api/modules/users/repositories/usersRepository"
)

func DeleteUserUseCase(id int) error {
	_, err := usersRepository.GetUserByID(id)

	if err != nil {
		return errors.New("user not found")
	}

	usersRepository.DeleteUser(id)
	return nil
}
