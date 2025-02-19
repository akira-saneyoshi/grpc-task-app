package di

import (
	"time"

	"github.com/akira-saneyoshi/task-app/application/handler"
	"github.com/akira-saneyoshi/task-app/application/usecase"
	"github.com/akira-saneyoshi/task-app/domain/service"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/model/db"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/sqlc"
	"github.com/akira-saneyoshi/task-app/utils/auth"
	"github.com/akira-saneyoshi/task-app/utils/clock"
	"github.com/akira-saneyoshi/task-app/utils/contextkey"
	"github.com/akira-saneyoshi/task-app/utils/identification"
)

func InitUser(qry db.Querier) *handler.UserHandler {
	repo := sqlc.NewSQLCUserRepository(qry)
	srv := service.NewUserService(repo)
	uc := usecase.NewUserUsecase(srv)
	return handler.NewUserHandler(uc)
}

func InitTask(qry db.Querier) *handler.TaskHandler {
	im := identification.NewUUIDManager()
	cm := clock.NewClockManager()
	cr := contextkey.NewContextReader()
	repo := sqlc.NewSQLCTaskRepository(qry)
	srv := service.NewTaskService(repo, im, cm)
	uc := usecase.NewTaskUsecase(srv)
	return handler.NewTaskHandler(uc, cr)
}

func InitAuth(issuer string, keyPath string, qry db.Querier, timeout time.Duration) (*handler.AuthHandler, error) {
	tm, err := auth.NewTokenManager(issuer, keyPath)
	if err != nil {
		return nil, err
	}
	repo := sqlc.NewSQLCUserRepository(qry)
	uc := usecase.NewAuthUsecase(repo, tm, timeout)
	return handler.NewAuthHandler(uc), nil
}
