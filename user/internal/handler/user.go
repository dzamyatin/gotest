package handler

import (
	api "app/user/api/user/proto"
	"app/user/internal/di/interceptor"
	use_case2 "app/user/internal/use_case"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	api.UnimplementedUserServer
	getAllUsersUseCase *use_case2.GetAllUsersUseCase
	getUserUseCase     *use_case2.GetUserUseCase
	createUserUseCase  *use_case2.CreateUserUseCase
}

func NewUserServer(
	getAllUsersUseCase *use_case2.GetAllUsersUseCase,
	getUserUseCase *use_case2.GetUserUseCase,
	createUserUseCase *use_case2.CreateUserUseCase,
) UserServer {
	return UserServer{
		getAllUsersUseCase: getAllUsersUseCase,
		getUserUseCase:     getUserUseCase,
		createUserUseCase:  createUserUseCase,
	}
}

func (s UserServer) Get(context context.Context, request *api.UserGetRequest) (*api.UserGetResponse, error) {
	user, err := s.getUserUseCase.Exec(use_case2.GetUserInput{
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

	user, err := s.createUserUseCase.Exec(
		use_case2.CreateUserInput{
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
	users, err := s.getAllUsersUseCase.Exec()

	if err != nil {
		return nil,
			status.Errorf(codes.Internal, "Internal error: %w", err)
	}

	for _, v := range users {
		response.Users = append(response.Users, &api.UserGetResponse{
			Uid:   v.Uid.String(),
			Login: v.Login,
		})
	}

	return &response, nil
}

func (s UserServer) Update(ctx context.Context, req *api.UserUpdateRequest) (*api.UserUpdateResponse, error) {
	ses := interceptor.GetSession(ctx)
	ses.NewGormSession()

	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
