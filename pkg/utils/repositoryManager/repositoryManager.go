package utils

import "os"

func GetRepository[T interface{}](repo1 T, repo2 T) T {
	env := os.Getenv("GO_APP_ENV")

	if env == "prod" {
		return repo1
	}

	return repo2
}
