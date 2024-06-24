package interceptor

import (
	"app/user/internal/di/session"
	"context"
	"google.golang.org/grpc"
	"log"
)

type SessionKey string

const (
	sessionKey = SessionKey("sessionKey")
)

func SessionServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	ses := session.NewSession(ctx)

	resp, err = handler(
		context.WithValue(ctx, sessionKey, ses),
		req,
	)

	return
}

func GetSession(ctx context.Context) *session.Session {
	ses, ok := ctx.Value(sessionKey).(*session.Session)

	if !ok {
		log.Fatalf("There is not session value to exclude")
	}

	return ses
}
