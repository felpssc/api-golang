package usecase

import (
	"errors"

	usersRepository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository"
)

type DeleteUserUseCase struct {
	usersRepository usersRepository.UsersRepository
}

func NewDeleteUserUseCase(usersRepository usersRepository.UsersRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		usersRepository: usersRepository,
	}
}

func (s DeleteUserUseCase) execute(id int) error {
	_, err := s.usersRepository.GetUserByID(id)

	if err != nil {
		return errors.New("user not found")
	}

	s.usersRepository.DeleteUser(id)
	return nil
}
