package repository

import entity "github.com/felpssc/api-golang/internal/modules/users/entities/user"

type UsersRepository interface {
	GetUsers() []entity.User
	GetUserByID(id int) (entity.User, error)
	CreateUser(name string, email string) (entity.User, error)
	FindUserByEmail(email string) (entity.User, error)
	UpdateUser(id int, name string) (entity.User, error)
	DeleteUser(id int) error
}
