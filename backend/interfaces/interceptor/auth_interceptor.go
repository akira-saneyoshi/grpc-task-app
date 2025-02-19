package interceptor

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/akira-saneyoshi/task-app/utils/auth"
	"github.com/akira-saneyoshi/task-app/utils/contextkey"
)

func NewAuthInterceptor(issuer string, keyPath string) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			token := req.Header().Get("Authorization")
			if token == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("[ERROR] invalid token"))
			}
			token = strings.TrimPrefix(token, "Bearer")
			token = strings.TrimSpace(token)

			tm, err := auth.NewTokenManager(issuer, keyPath)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}
			uid, err := tm.GetUserID(token)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			cw := contextkey.NewContextWriter()
			ctx = cw.SetUserID(ctx, uid)

			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
