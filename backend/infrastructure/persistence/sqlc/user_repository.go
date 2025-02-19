package sqlc

import (
	"context"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	v_user "github.com/akira-saneyoshi/task-app/domain/object/value/v_user"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/model/db"
	"github.com/akira-saneyoshi/task-app/utils/convert"
)

// SQLC implementation of user persistence.
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
	return &entity.User{
		ID:        value.NewID(res.ID),
		Name:      v_user.NewName(res.Name),
		Email:     v_user.NewEmail(res.Email),
		Password:  v_user.NewPassword(res.Password),
		IsActive:  res.IsActive,
		CreatedAt: convert.ConvertNullTime(res.CreatedAt),
		UpdatedAt: convert.ConvertNullTime(res.UpdatedAt),
	}, nil
}

func (r *SQLCUserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	res, err := r.Querier.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:        value.NewID(res.ID),
		Name:      v_user.NewName(res.Name),
		Email:     v_user.NewEmail(res.Email),
		Password:  v_user.NewPassword(res.Password),
		IsActive:  res.IsActive,
		CreatedAt: convert.ConvertNullTime(res.CreatedAt),
		UpdatedAt: convert.ConvertNullTime(res.UpdatedAt),
	}, nil
}
