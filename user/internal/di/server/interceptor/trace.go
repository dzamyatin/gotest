package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func InterceptorTrace(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

	resp, err = handler(ctx, req)

	return
}
