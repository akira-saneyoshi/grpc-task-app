package repository

import (
	"context"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
)

// Persist UserEntity
type IUserRepository interface {
	FindUserByID(ctx context.Context, id string) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
