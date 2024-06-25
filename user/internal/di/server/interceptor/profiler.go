package interceptor

import (
	"app/user/internal/di/static"
	"context"
	"google.golang.org/grpc"
)

func InterceptorProfiler(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	//dont know what the point to lablel context. dont see a result
	ctx = static.GetProfiler().LabelCtxt(
		ctx,
		"grpc",
		info.FullMethod,
	)

	resp, err = handler(ctx, req)

	return
}