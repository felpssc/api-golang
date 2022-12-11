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

func TestListUsers(t *testing.T) {
	usersRepository := repository.NewLocalUsersRepository()

	usersRepository.CreateUser(
		"John Doe",
		"johndoe@email.com",
		"123456",
	)

	usersRepository.CreateUser(
		"Jane Doe",
		"janedoe@email.com",
		"123456",
	)

	listUsersUseCase := NewListUsersUseCase(usersRepository)

	users := listUsersUseCase.execute()

	assert.Equal(t, 2, len(users))
	assert.Equal(t, "John Doe", users[0].Name)
	assert.Equal(t, "Jane Doe", users[1].Name)
}
