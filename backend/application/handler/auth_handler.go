package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/akira-saneyoshi/task-app/application"
	"github.com/akira-saneyoshi/task-app/application/usecase"
	"github.com/akira-saneyoshi/task-app/domain"
	dto_auth "github.com/akira-saneyoshi/task-app/interfaces/dto/dto_auth"
	auth_v1 "github.com/akira-saneyoshi/task-app/interfaces/proto/auth/v1"
)

type AuthHandler struct {
	usecase.IAuthUsecase
}

func NewAuthHandler(uc usecase.IAuthUsecase) *AuthHandler {
	return &AuthHandler{uc}
}

func (h *AuthHandler) Login(ctx context.Context, arg *connect.Request[auth_v1.LoginRequest]) (*connect.Response[auth_v1.LoginResponse], error) {
	params := dto_auth.NewLoginParams(arg.Msg.Email, arg.Msg.Password)
	info, err := h.IAuthUsecase.Login(ctx, params)
	if err != nil {
		switch e := err.(type) {
		case *application.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *application.ErrLoginFailed:
			return nil, connect.NewError(connect.CodeUnauthenticated, e)
		case *domain.ErrNotFound:
			return nil, connect.NewError(connect.CodeUnauthenticated, e)
		case *application.ErrInternal:
			return nil, connect.NewError(connect.CodeInternal, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	return connect.NewResponse(&auth_v1.LoginResponse{
		Token: info.Token(),
	}), nil
}
