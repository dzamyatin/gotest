package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func InterceptorProfiler(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	//prof := static.GetProfiler()

	resp, err = handler(ctx, req)

	return
}
