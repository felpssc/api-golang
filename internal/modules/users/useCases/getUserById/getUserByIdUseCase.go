package usecase

import (
	user "github.com/felpssc/api-golang/internal/modules/users/entities/user"
	usersRepository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository"
)

type GetUserByIdUseCase struct {
	usersRepository usersRepository.UsersRepository
}

func NewGetUserByIdUseCase(usersRepository usersRepository.UsersRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		usersRepository: usersRepository,
	}
}

func (s GetUserByIdUseCase) execute(id int) (user.User, error) {
	user, err := s.usersRepository.GetUserByID(id)

	if err != nil {
		return user, err
	}

	return user, nil
}
