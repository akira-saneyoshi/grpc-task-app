package service

import (
	"context"

	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	"github.com/akira-saneyoshi/task-app/domain/repository"
)

type IUserService interface {
	FindUserByID(ctx context.Context, id string) (*entity.User, error)
}

type UserService struct {
	repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) FindUserByID(ctx context.Context, id string) (*entity.User, error) {
	if err := value.NewID(id).Validate(); err != nil {
		return nil, err
	}
	user, err := s.IUserRepository.FindUserByID(ctx, id)
	if err != nil {
		return nil, &domain.ErrNotFound{Msg: "user not found"}
	}
	return user, nil
}
