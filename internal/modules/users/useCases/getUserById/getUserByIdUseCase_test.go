package usecase

import (
	"os"
	"testing"

	repository "github.com/felpssc/api-golang/internal/modules/users/repositories/usersRepository/implementations"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_APP_ENV", "dev")
}

func TestGetUserById(t *testing.T) {
	usersRepository := repository.NewLocalUsersRepository()

	createdUser, _ := usersRepository.CreateUser(
		"Teste",
		"teste@email.com",
		"123456",
	)

	getUserByIdUseCase := NewGetUserByIdUseCase(usersRepository)

	user, err := getUserByIdUseCase.execute(createdUser.ID)

	assert.Nil(t, err)
	assert.Equal(t, createdUser.ID, user.ID)
	assert.Equal(t, createdUser.Name, user.Name)
	assert.Equal(t, createdUser.Email, user.Email)
}
