package usersRepository

import (
	"errors"

	user "github.com/felpssc/api-golang/api/modules/users/entities/user"
)

var users = []user.User{
	{ID: 1, Name: "John", Email: "Laney.Hudson@hotmail.com", IsAdmin: true},
	{ID: 2, Name: "Jane", Email: "Laney.Hudson@hotmail.com", IsAdmin: false},
}

func GetUsers() []user.User {
	return users
}

func GetUserByID(id int) (user.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return user.User{}, errors.New("user not found")
}

func CreateUser(name string, email string) (user.User, error) {
	newUser := user.User{
		ID:      len(users) + 1,
		Name:    name,
		Email:   email,
		IsAdmin: false,
	}

	users = append(users, newUser)

	return newUser, nil
}

func FindUserByEmail(email string) (user.User, error) {

	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}

	return user.User{}, errors.New("user not found")
}

func UpdateUser(id int, name string) (user.User, error) {
	user, err := GetUserByID(id)

	if err != nil {
		return user, errors.New("user not found")
	}

	user.Name = name

	var userIndex int

	for i, u := range users {
		if u.ID == id {
			userIndex = i
			break
		}
	}

	users[userIndex] = user

	return user, nil
}

func DeleteUser(id int) error {
	var userIndex int

	for i, u := range users {
		if u.ID == id {
			userIndex = i
			break
		}
	}

	users = append(users[:userIndex], users[userIndex+1:]...)

	return nil
}
