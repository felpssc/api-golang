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

func TestDeleteUser(t *testing.T) {
	usersRepository := repository.NewLocalUsersRepository()

	usersRepository.CreateUser(
		"John Doe",
		"johndoe@email.com",
		"123456",
	)

	deleteUserUseCase := NewDeleteUserUseCase(usersRepository)

	err := deleteUserUseCase.execute(1)

	assert.Nil(t, err)

	user, _ := usersRepository.GetUserByID(1)

	assert.Equal(t, 0, user.ID)
}
