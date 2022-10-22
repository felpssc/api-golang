package usecase

import (
	user "github.com/felpssc/api-golang/internal/modules/users/entities/user"
	usersRepository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository"
)

type ListUsersUseCase struct {
	usersRepository usersRepository.UsersRepository
}

func NewListUsersUseCase(usersRepository usersRepository.UsersRepository) *ListUsersUseCase {
	return &ListUsersUseCase{
		usersRepository: usersRepository,
	}
}

func (s ListUsersUseCase) execute() []user.User {
	users := s.usersRepository.GetUsers()
	return users
}
