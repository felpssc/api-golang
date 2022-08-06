package createUserUseCase

import (
	"errors"

	user "github.com/felpssc/api-golang/api/modules/users/entities/user"
	usersRepository "github.com/felpssc/api-golang/api/modules/users/repositories/usersRepository"
)

func CreateUserUseCase(name string, email string) (user.User, error) {

	userAlreadyExists, _ := usersRepository.FindUserByEmail(email)

	if userAlreadyExists.ID > 0 {
		return user.User{}, errors.New("user already exists")
	}

	createdUser, _ := usersRepository.CreateUser(name, email)

	return createdUser, nil
}
