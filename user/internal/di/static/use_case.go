package static

import (
	"app/user/internal/use_case"
)

func GetCreateUserUseCase() *use_case.CreateUserUseCase {
	object := use_case.NewCreateUserUseCase(GetUserRepository())
	return &object
}

func GetAllUsersUseCase() *use_case.GetAllUsersUseCase {
	object := use_case.NewGetAllUsersUseCase(GetUserRepository())
	return &object
}

func GetUserUseCase() *use_case.GetUserUseCase {
	object := use_case.NewGetUserUseCase(GetUserRepository())
	return &object
}
