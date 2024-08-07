package static

import (
	"app/user/internal/repository"
)

func GetUserRepository() repository.UserRepositoryInterface {
	return repository.NewUserRepository(GetDB())
}
