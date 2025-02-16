package sqlc

import (
	"testing"

	"github.com/akira-saneyoshi/task-app/domain/repository"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/sqlc"
)

func TestUserRepository_NewUserRepository(tt *testing.T) {
	tt.Run("異常系: structがinterfaceを実装しているか", func(t *testing.T) {
		var _ repository.IUserRepository = (*sqlc.SQLCUserRepository)(nil)
	})

}
