package usecase

import (
	"errors"

	user "github.com/felpssc/api-golang/internal/modules/users/entities/user"
	repository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository"
)

type CreateUserUseCase struct {
	UsersRepository repository.UsersRepository
}

func NewCreateUserUseCase(usersRepository repository.UsersRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UsersRepository: usersRepository,
	}
}

func (s CreateUserUseCase) execute(name string, email string) (user.User, error) {

	userAlreadyExists, _ := s.UsersRepository.FindUserByEmail(email)

	if userAlreadyExists.ID > 0 {
		return user.User{}, errors.New("user already exists")
	}

	createdUser, _ := s.UsersRepository.CreateUser(name, email)

	return createdUser, nil
}
