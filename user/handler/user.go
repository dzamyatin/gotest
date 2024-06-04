package handler

import (
	api "app/user/api/user/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	api.UnimplementedUserServer
}

func (s UserServer) Create(context.Context, *api.UserCreateRequest) (*api.UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
