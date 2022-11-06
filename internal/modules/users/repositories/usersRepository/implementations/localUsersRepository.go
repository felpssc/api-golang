package repository

import (
	entity "github.com/felpssc/api-golang/internal/modules/users/entities/user"
)

var users = []entity.User{
	{
		ID:       1,
		Name:     "John",
		Email:    "Laney.Hudson@hotmail.com",
		Password: "e8ene83nf94n2d23a1o3",
		IsAdmin:  true,
	},
	{
		ID:       2,
		Name:     "Jane",
		Email:    "Laney.Hudson@hotmail.com",
		Password: "8r3e2dnfn3e99fa220d",
		IsAdmin:  false,
	},
}

type LocalUsersRepository struct{}

func NewLocalUsersRepository() *LocalUsersRepository {
	return &LocalUsersRepository{}
}

func (u LocalUsersRepository) CreateUser(name string, email string, password string) (entity.User, error) {
	user := entity.User{
		ID:      len(users) + 1,
		Name:    name,
		Email:   email,
		IsAdmin: false,
	}

	users = append(users, user)

	return user, nil
}

func (u LocalUsersRepository) FindUserByEmail(email string) (entity.User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}

	return entity.User{}, nil
}

func (u LocalUsersRepository) GetUsers() []entity.User {
	return users
}

func (u LocalUsersRepository) GetUserByID(id int) (entity.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return entity.User{}, nil
}

func (u LocalUsersRepository) UpdateUser(id int, name string) (entity.User, error) {
	for index, user := range users {
		if user.ID == id {
			users[index].Name = name
			return users[index], nil
		}
	}

	return entity.User{}, nil
}

func (u LocalUsersRepository) DeleteUser(id int) error {
	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			return nil
		}
	}

	return nil
}