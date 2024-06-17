package config

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

const (
	sessionKey = "sessionKey"
)

func SessionServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	session := NewSession(ctx)

	context.WithValue(ctx, sessionKey, session)
	resp, err = handler(ctx, req)

	return
}

func GetSession(ctx context.Context) *Session {
	session, ok := ctx.Value(sessionKey).(Session)

	if !ok {
		log.Fatalf("There is not session value to exclude")
	}

	return &session
}
