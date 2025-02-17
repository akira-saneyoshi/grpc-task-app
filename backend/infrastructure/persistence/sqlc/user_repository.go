package sqlc

import (
	"context"
	"time"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	user "github.com/akira-saneyoshi/task-app/domain/object/value/user"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/model/db"
)

type SQLCUserRepository struct {
	db.Querier
}

func NewSQLCUserRepository(qry db.Querier) *SQLCUserRepository {
	return &SQLCUserRepository{qry}
}

func (r *SQLCUserRepository) FindUserByID(ctx context.Context, id string) (*entity.User, error) {
	res, err := r.Querier.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var createdAt time.Time
	if res.CreatedAt.Valid {
		createdAt = res.CreatedAt.Time
	}

	var updatedAt time.Time
	if res.UpdatedAt.Valid {
		updatedAt = res.UpdatedAt.Time
	}

	return &entity.User{
		ID:        value.NewID(res.ID),
		Name:      user.NewName(res.Name),
		Email:     user.NewEmail(res.Email),
		Password:  user.NewPassword(res.Password),
		IsActive:  res.IsActive,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (r *SQLCUserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	res, err := r.Querier.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	var createdAt time.Time
	if res.CreatedAt.Valid {
		createdAt = res.CreatedAt.Time
	}

	var updatedAt time.Time
	if res.UpdatedAt.Valid {
		updatedAt = res.UpdatedAt.Time
	}

	return &entity.User{
		ID:        value.NewID(res.ID),
		Name:      user.NewName(res.Name),
		Email:     user.NewEmail(res.Email),
		Password:  user.NewPassword(res.Password),
		IsActive:  res.IsActive,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
