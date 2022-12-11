package usecase

import (
	"os"
	"testing"

	"github.com/bxcodec/faker/v3"
	user "github.com/felpssc/api-golang/internal/modules/users/entities/user"
	repository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository/implementations"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_APP_ENV", "dev")
}

func TestCreateUser(t *testing.T) {
	usersRepository := repository.NewLocalUsersRepository()

	createUserUseCase := NewCreateUserUseCase(usersRepository)

	user := user.User{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	createdUser, err := createUserUseCase.execute(user.Name, user.Email, user.Password)

	assert.Nil(t, err)

	assert.Equal(t, user.Name, createdUser.Name)

	assert.NotEqual(t, 0, createdUser.ID)
}

func TestIfEmailExists(t *testing.T) {
	usersRepository := repository.NewLocalUsersRepository()

	createUserUseCase := NewCreateUserUseCase(usersRepository)

	user := user.User{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	createdUser, err := createUserUseCase.execute(user.Name, user.Email, user.Password)

	assert.Nil(t, err)

	assert.Equal(t, user.Name, createdUser.Name)

	assert.NotEqual(t, 0, createdUser.ID)

	_, err = createUserUseCase.execute(user.Name, user.Email, user.Password)

	assert.NotNil(t, err)

	assert.Equal(t, "user already exists", err.Error())
}
