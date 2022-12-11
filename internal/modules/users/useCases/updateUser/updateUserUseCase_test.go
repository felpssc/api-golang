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

func TestUpdateUser(t *testing.T) {
	usersRepository := repository.NewLocalUsersRepository()

	usersRepository.CreateUser(
		"John Doe",
		"johndoe@email.com",
		"123456",
	)

	updateUserUseCase := NewUpdateUserUseCase(usersRepository)

	user, err := updateUserUseCase.execute(1, "Jane Doe")

	assert.Nil(t, err)

	assert.Equal(t, "Jane Doe", user.Name)
}
