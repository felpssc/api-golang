package getUserById

import (
	user "github.com/felpssc/api-golang/api/modules/users/entities/user"
	usersRepository "github.com/felpssc/api-golang/api/modules/users/repositories/usersRepository"
)

func GetUserByIdUseCase(id int) (user.User, error) {
	user, err := usersRepository.GetUserByID(id)

	if err != nil {
		return user, err
	}

	return user, nil
}
