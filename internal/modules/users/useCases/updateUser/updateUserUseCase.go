package usecase

import (
	"errors"

	user "github.com/felpssc/api-golang/internal/modules/users/entities/user"
	usersRepository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository"
)

type UpdateUserUseCase struct {
	usersRepository usersRepository.UsersRepository
}

func NewUpdateUserUseCase(usersRepository usersRepository.UsersRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		usersRepository: usersRepository,
	}
}

func (s UpdateUserUseCase) execute(id int, name string) (user.User, error) {
	user, err := s.usersRepository.GetUserByID(id)

	if err != nil {
		return user, errors.New("user not found")
	}

	updatedUser, err := s.usersRepository.UpdateUser(id, name)

	if err != nil {
		return user, err
	}

	return updatedUser, nil
}
