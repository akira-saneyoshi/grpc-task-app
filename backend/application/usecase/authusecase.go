package usecase

import (
	"context"
	"time"

	"github.com/akira-saneyoshi/task-app/application"
	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/domain/repository"
	"github.com/akira-saneyoshi/task-app/interfaces/dto"
	"github.com/akira-saneyoshi/task-app/utils/auth"
	"golang.org/x/crypto/bcrypt"
)

type IAuthUsecase interface {
	Login(ctx context.Context, arg *dto.LoginParams) (*dto.UserInfo, error)
}

type AuthUsecase struct {
	repository.IUserRepository
	auth.ITokenManager
	timeout time.Duration
}

func NewAuthUsecase(repo repository.IUserRepository, tm auth.ITokenManager, timeout time.Duration) *AuthUsecase {
	return &AuthUsecase{repo, tm, timeout}
}

func (u *AuthUsecase) Login(ctx context.Context, arg *dto.LoginParams) (*dto.UserInfo, error) {
	if err := arg.Validate(); err != nil {
		return nil, err
	}
	user, err := u.IUserRepository.FindUserByEmail(ctx, arg.Email())
	if err != nil {
		return nil, &domain.ErrNotFound{Msg: "[ERROR] user not found"}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.Value()), []byte(arg.Password())); err != nil {
		return nil, &application.ErrLoginFailed{Msg: "[ERROR] password does not match"}
	}
	token, err := u.ITokenManager.CreateToken(user.ID.Value(), u.timeout)
	if err != nil {
		return nil, &application.ErrInternal{Msg: "[ERROR] failed to create token"}
	}
	return dto.NewUserInfo(user.ID.Value(), user.Email.Value(), token), nil
}
