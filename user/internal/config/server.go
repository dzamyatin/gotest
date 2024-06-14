package config

import "app/user/internal/handler"

func GetUserServer() *handler.UserServer {
	return syncGetOrCreateByType(
		func() *handler.UserServer {
			obj := handler.NewUserServer(
				GetAllUsersUseCase(),
				GetUserUseCase(),
				GetCreateUserUseCase(),
			)

			return &obj
		},
	)
}
