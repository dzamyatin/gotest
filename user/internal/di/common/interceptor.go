package common

import (
	"app/user/internal/di/session"
	"context"
	"google.golang.org/grpc"
	"log"
)

const (
	sessionKey = "sessionKey"
)

func SessionServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	ses := session.NewSession(ctx)

	context.WithValue(ctx, sessionKey, ses)
	resp, err = handler(ctx, req)

	return
}

func GetSession(ctx context.Context) *session.Session {
	ses, ok := ctx.Value(sessionKey).(session.Session)

	if !ok {
		log.Fatalf("There is not session value to exclude")
	}

	return &ses
}
