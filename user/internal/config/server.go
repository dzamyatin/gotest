package config

import (
	"app/user/internal/handler"
)

func GetUserServer() *handler.UserServer {
	obj := handler.NewUserServer(
		GetAllUsersUseCase(),
		GetUserUseCase(),
		GetCreateUserUseCase(),
	)

	return &obj
}
