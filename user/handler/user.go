package handler

import (
	api "app/user/api/user/proto"
	"app/user/lib"
	"app/user/use_case"
	"context"
)

type UserServer struct {
	api.UnimplementedUserServer
}

func (s UserServer) Create(context context.Context, request *api.UserCreateRequest) (*api.UserCreateResponse, error) {

	lib.EventBusInstance.Dispatch(use_case.CreateUserInput{
		Login: request.Login,
	})

	//use_case.InitCreateUserUseCase().Execute(use_case.CreateUserInput{})

	return nil, nil
}
