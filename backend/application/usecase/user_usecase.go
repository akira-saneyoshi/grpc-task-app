package usecase

import (
	"context"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/service"
	"github.com/akira-saneyoshi/task-app/interfaces/dto"
)

type IUserUsecase interface {
	FindUserByID(ctx context.Context, id *dto.IDParam) (*entity.User, error)
}

type UserUsecase struct {
	service.IUserService
}

func NewUserUsecase(srv service.IUserService) *UserUsecase {
	return &UserUsecase{srv}
}

func (u *UserUsecase) FindUserByID(ctx context.Context, id *dto.IDParam) (*entity.User, error) {
	if err := id.Validate(); err != nil {
		return nil, err
	}
	return u.IUserService.FindUserByID(ctx, id.Value())
}
