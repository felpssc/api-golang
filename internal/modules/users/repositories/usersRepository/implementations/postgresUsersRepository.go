package repository

import (
	entity "github.com/felpssc/api-golang/internal/modules/users/entities/user"
	config "github.com/felpssc/api-golang/pkg/config"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type PostgresUsersRepository struct{}

func NewPostgresUsersRepository() *PostgresUsersRepository {
	return &PostgresUsersRepository{}
}

func (u PostgresUsersRepository) CreateUser(name string, email string, password string) (entity.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	config.GetDB().Create(&user)

	return user, nil
}

func (u PostgresUsersRepository) FindUserByEmail(email string) (entity.User, error) {
	var user entity.User

	config.GetDB().Where("email = ?", email).First(&user)

	return user, nil
}

func (u PostgresUsersRepository) GetUsers() []entity.User {
	var users []entity.User

	config.GetDB().Find(&users)

	return users
}

func (u PostgresUsersRepository) GetUserByID(id int) (entity.User, error) {
	var user entity.User

	config.GetDB().Where("id = ?", id).First(&user)

	return user, nil
}

func (u PostgresUsersRepository) UpdateUser(id int, name string) (entity.User, error) {
	var user entity.User

	config.GetDB().Where("id = ?", id).First(&user)

	user.Name = name

	config.GetDB().Save(&user)

	return user, nil
}

func (u PostgresUsersRepository) DeleteUser(id int) error {
	var user entity.User

	config.GetDB().Where("id = ?", id).First(&user)

	config.GetDB().Delete(&user)

	return nil
}
