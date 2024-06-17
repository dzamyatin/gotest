package server

import (
	"app/user/internal/di/static"
	"app/user/internal/handler"
)

func GetUserServer() *handler.UserServer {
	obj := handler.NewUserServer(
		static.GetAllUsersUseCase(),
		static.GetUserUseCase(),
		static.GetCreateUserUseCase(),
	)

	return &obj
}
