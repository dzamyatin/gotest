package handler

import (
	api "app/user/api/user/proto"
	"app/user/use_case"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	api.UnimplementedUserServer
}

func (s UserServer) Create(context context.Context, request *api.UserCreateRequest) (*api.UserCreateResponse, error) {
	//lib.EventBusInstance.Dispatch(use_case.CreateUserInput{
	//	Login: request.Login,
	//})

	user, err := use_case.InitCreateUserUseCase().Exec(
		use_case.CreateUserInput{
			Login: request.Login,
		},
	)

	if err != nil {
		return nil,
			status.Errorf(codes.Unimplemented, "method Create not implemented")
	}

	return &api.UserCreateResponse{
		Uid: user.Uid(),
	}, nil
}
