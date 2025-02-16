package service

import (
	"context"
	"testing"
	"time"

	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	user "github.com/akira-saneyoshi/task-app/domain/object/value/user"
	"github.com/akira-saneyoshi/task-app/domain/service"
	"github.com/akira-saneyoshi/task-app/test/mocks"

	"github.com/stretchr/testify/require"
)

func TestUserService_NewUserService(tt *testing.T) {
	tt.Run("異常系: structがinterfaceを実装しているか", func(t *testing.T) {
		var _ service.IUserService = (*service.UserService)(nil)
	})
}

func TestUserService_FindUserByID(tt *testing.T) {
	now := time.Now().UTC()
	ctx := context.Background()
	id := "id"
	user := &entity.User{
		ID:        value.NewID(id),
		Name:      user.NewName("test-user"),
		Email:     user.NewEmail("email"),
		Password:  user.NewPassword("pass"),
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	tt.Run("正常系: 正しい入力の場合", func(t *testing.T) {
		repo := new(mocks.IUserRepository)
		repo.On("FindUserByID", ctx, id).Return(user, nil)
		srv := service.NewUserService(repo)
		ret, err := srv.FindUserByID(ctx, id)

		require.NoError(t, err, "エラーが発生しないこと")
		require.Equal(t, id, ret.ID.Value())
		require.Equal(t, user.Name.Value(), ret.Name.Value())
		require.Equal(t, user.Email.Value(), ret.Email.Value())
		require.Equal(t, user.Password.Value(), ret.Password.Value())
		require.Equal(t, user.IsActive, ret.IsActive)
		require.Equal(t, user.CreatedAt, ret.CreatedAt)
		require.Equal(t, user.UpdatedAt, ret.UpdatedAt)
		repo.AssertExpectations(t)
	})
	tt.Run("準正常系: UserIDが空の場合", func(t *testing.T) {
		errExp := &domain.ErrValidationFailed{Msg: "id is empty"}
		id := ""
		repo := new(mocks.IUserRepository)
		srv := service.NewUserService(repo)
		_, err := srv.FindUserByID(ctx, id)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
	})
	tt.Run("準正常系: 存在しないUserIDの場合", func(t *testing.T) {
		errExp := &domain.ErrNotFound{Msg: "user not found"}
		id := "another"
		repo := new(mocks.IUserRepository)
		repo.On("FindUserByID", ctx, id).Return(nil, errExp)
		srv := service.NewUserService(repo)
		_, err := srv.FindUserByID(ctx, id)

		require.EqualError(t, err, errExp.Error(), "エラーが一致すること")
		repo.AssertExpectations(t)
	})
}
