package di

import (
	"app/user/internal/config"
	"context"
	"google.golang.org/grpc"
	"log"
)

const (
	sessionKey = "sessionKey"
)

func SessionServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	session := config.NewSession(ctx)

	context.WithValue(ctx, sessionKey, session)
	resp, err = handler(ctx, req)

	return
}

func GetSession(ctx context.Context) *config.Session {
	session, ok := ctx.Value(sessionKey).(config.Session)

	if !ok {
		log.Fatalf("There is not session value to exclude")
	}

	return &session
}
