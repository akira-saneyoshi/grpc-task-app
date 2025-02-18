package usecase

import (
	"context"
	"html"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/service"
	"github.com/akira-saneyoshi/task-app/interfaces/dto"
)

type ITaskUsecase interface {
	FindTasksByUserID(ctx context.Context, userID *dto.IDParam) ([]*entity.Task, error)
	CreateTask(ctx context.Context, arg *dto.CreateTaskParams) (string, error)
	UpdateTaskDetails(ctx context.Context, arg *dto.UpdateTaskDetailsParams) error
	UpdateTaskDueDate(ctx context.Context, arg *dto.UpdateTaskDueDateParams) error
	UpdateTaskStatus(ctx context.Context, arg *dto.UpdateTaskStatusParams) error
	DeleteTask(ctx context.Context, id *dto.IDParam, userID *dto.IDParam) error
}

type TaskUsecase struct {
	service.ITaskService
}

func NewTaskUsecase(srv service.ITaskService) *TaskUsecase {
	return &TaskUsecase{srv}
}

func (u *TaskUsecase) FindTasksByUserID(ctx context.Context, userID *dto.IDParam) ([]*entity.Task, error) {
	if err := userID.Validate(); err != nil {
		return nil, err
	}
	return u.ITaskService.FindTasksByUserID(ctx, userID.Value())
}

func (u *TaskUsecase) CreateTask(ctx context.Context, arg *dto.CreateTaskParams) (string, error) {
	if err := arg.Validate(); err != nil {
		return "", err
	}
	return u.ITaskService.CreateTask(
		ctx,
		arg.UserID(),
		html.EscapeString(arg.Title()),
		arg.Description(),
		string(arg.Status()),
		arg.DueDate(),
	)
}

func (u *TaskUsecase) UpdateTaskDetails(ctx context.Context, arg *dto.UpdateTaskDetailsParams) error {
	if err := arg.Validate(); err != nil {
		return err
	}
	return u.ITaskService.UpdateTaskDetails(
		ctx,
		arg.ID(),
		arg.UserID(),
		html.EscapeString(arg.Title()),
		arg.Description(),
		string(arg.Status()),
		arg.DueDate(),
	)
}

func (u *TaskUsecase) UpdateTaskDueDate(ctx context.Context, arg *dto.UpdateTaskDueDateParams) error {
	if err := arg.Validate(); err != nil {
		return err
	}
	return u.ITaskService.UpdateTaskDueDate(
		ctx,
		arg.ID(),
		arg.UserID(),
		arg.DueDate(),
	)
}

func (u *TaskUsecase) UpdateTaskStatus(ctx context.Context, arg *dto.UpdateTaskStatusParams) error {
	if err := arg.Validate(); err != nil {
		return err
	}
	return u.ITaskService.UpdateTaskStatus(
		ctx,
		arg.ID(),
		arg.UserID(),
		string(arg.Status()),
	)
}

func (u *TaskUsecase) DeleteTask(ctx context.Context, id *dto.IDParam, userID *dto.IDParam) error {
	if err := id.Validate(); err != nil {
		return err
	}
	if err := userID.Validate(); err != nil {
		return err
	}
	return u.ITaskService.DeleteTask(ctx, id.Value(), userID.Value())
}
