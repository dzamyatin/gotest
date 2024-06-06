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

func (s UserServer) Get(context context.Context, request *api.UserGetRequest) (*api.UserGetResponse, error) {
	user, err := use_case.InitGetUserUseCase().Exec(use_case.GetUserInput{
		Uid: request.Uid,
	})

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "There is no such user")
	}

	return &api.UserGetResponse{Uid: user.Uid.String(), Login: user.Login}, nil
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
			status.Errorf(codes.Internal, "Not found")
	}

	return &api.UserCreateResponse{
		Uid: user.Uid.String(),
	}, nil
}

func (s UserServer) GetAll(context context.Context, request *api.UserGetAllRequest) (*api.UserGetAllResponse, error) {
	response := api.UserGetAllResponse{}
	users, err := use_case.InitGetAllUsersUseCase().Exec()

	if err != nil {
		return nil,
			status.Errorf(codes.Internal, "Internal error")
	}

	for _, v := range users {
		response.Users = append(response.Users, &api.UserGetResponse{
			Uid:   v.Uid.String(),
			Login: v.Login,
		})
	}

	return &response, nil
}
