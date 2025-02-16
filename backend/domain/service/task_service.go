package service

import (
	"context"
	"time"

	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	task "github.com/akira-saneyoshi/task-app/domain/object/value/task"
	"github.com/akira-saneyoshi/task-app/domain/repository"
	"github.com/akira-saneyoshi/task-app/utils/clock"
	"github.com/akira-saneyoshi/task-app/utils/identification"
)

// TaskService はタスクに関するドメインロジックを提供します。
type ITaskService interface {
	FindTaskByID(ctx context.Context, id string) (*entity.Task, error)
	FindTasksByUserID(ctx context.Context, userID string) ([]*entity.Task, error)
	CreateTask(ctx context.Context, userID string, title string, description *string, status string, dueDate *time.Time) (string, error)
	ChangeTaskDetails(ctx context.Context, id string, userID string, title string, description *string, status string, dueDate *time.Time) error // Title, Description, Status, DueDateをまとめて変更する関数に変更
	ChangeTaskStatus(ctx context.Context, id string, userID string, status string) error                                                         // Status変更用関数（そのまま残す）
	ChangeTaskDueDate(ctx context.Context, id string, userID string, dueDate *time.Time) error                                                   // DueDate変更用関数（そのまま残す）
	DeleteTask(ctx context.Context, id, userID string) error
}

type TaskService struct {
	repository.ITaskRepository
	identification.IIDManager
	clock.IClockManager
}

func NewTaskService(repo repository.ITaskRepository, idManager identification.IIDManager, clockManager clock.IClockManager) *TaskService {
	return &TaskService{repo, idManager, clockManager}
}

func (s *TaskService) FindTaskByID(ctx context.Context, id string) (*entity.Task, error) {
	if err := value.NewID(id).Validate(); err != nil {
		return nil, err
	}
	task, err := s.ITaskRepository.FindTaskByID(ctx, id)
	if err != nil {
		return nil, &domain.ErrNotFound{Msg: "task not found"}
	}
	return task, nil
}

func (s *TaskService) FindTasksByUserID(ctx context.Context, userID string) ([]*entity.Task, error) {
	if err := value.NewID(userID).Validate(); err != nil {
		return nil, err
	}
	tasks, err := s.ITaskRepository.FindTasksByUserID(ctx, userID)
	if err != nil {
		return nil, &domain.ErrQueryFailed{}
	}
	return tasks, nil
}

func (s *TaskService) CreateTask(ctx context.Context, userID string, title string, description *string, status string, dueDate *time.Time) (string, error) {
	now := s.IClockManager.GetNow()
	taskStatus := entity.StatusPending // デフォルトStatusをPendingに設定
	if status != "" {
		taskStatus = entity.Status(status) // 引数statusが指定されている場合はそちらを使う
	}

	arg := &entity.Task{
		ID:          value.NewID(s.IIDManager.GenerateID()),
		UserID:      value.NewID(userID),
		Title:       task.NewTitle(title),
		Description: description,
		Status:      taskStatus,
		DueDate:     dueDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := arg.Validate(); err != nil {
		return "", err
	}
	createdID, err := s.ITaskRepository.CreateTask(ctx, arg)
	if err != nil {
		return "", &domain.ErrQueryFailed{}
	}
	return createdID, nil
}

// ChangeTaskDetails はタスクのTitle, Description, Status, DueDateをまとめて変更します。
func (s *TaskService) ChangeTaskDetails(ctx context.Context, id string, userID string, title string, description *string, status string, dueDate *time.Time) error {
	if err := value.NewID(id).Validate(); err != nil {
		return err
	}
	if err := value.NewID(userID).Validate(); err != nil {
		return err
	}
	task, err := s.ITaskRepository.FindTaskByID(ctx, id)
	if err != nil {
		return &domain.ErrNotFound{Msg: "task not found"}
	}
	if !task.UserID.Equal(userID) {
		return &domain.ErrPermissionDenied{}
	}
	task.Title = title
	task.Description = description
	taskStatus := entity.Status(status)
	if status != "" && !entity.IsValidStatus(taskStatus) {
		return &domain.ErrValidationFailed{Msg: "invalid status value"}
	}
	if status != "" {
		task.Status = taskStatus
	}
	task.DueDate = dueDate
	task.UpdatedAt = s.IClockManager.GetNow()
	if err := task.Validate(); err != nil {
		return err
	}
	if err := s.ITaskRepository.UpdateTask(ctx, task); err != nil {
		return &domain.ErrQueryFailed{}
	}
	return nil
}

func (s *TaskService) ChangeTaskStatus(ctx context.Context, id string, userID string, status string) error {
	if err := value.NewID(id).Validate(); err != nil {
		return err
	}
	if err := value.NewID(userID).Validate(); err != nil {
		return err
	}
	task, err := s.ITaskRepository.FindTaskByID(ctx, id)
	if err != nil {
		return &domain.ErrNotFound{Msg: "task not found"}
	}
	if !task.UserID.Equal(userID) {
		return &domain.ErrPermissionDenied{}
	}
	taskStatus := entity.Status(status)
	if !entity.IsValidStatus(taskStatus) {
		return &domain.ErrValidationFailed{Msg: "invalid status value"}
	}
	task.Status = taskStatus
	task.UpdatedAt = s.IClockManager.GetNow()
	if err := s.ITaskRepository.UpdateTask(ctx, task); err != nil {
		return &domain.ErrQueryFailed{}
	}
	return nil
}

func (s *TaskService) ChangeTaskDueDate(ctx context.Context, id string, userID string, dueDate *time.Time) error { // DueDate変更用関数（そのまま残す）
	if err := value.NewID(id).Validate(); err != nil {
		return err
	}
	if err := value.NewID(userID).Validate(); err != nil {
		return err
	}
	task, err := s.ITaskRepository.FindTaskByID(ctx, id)
	if err != nil {
		return &domain.ErrNotFound{Msg: "task not found"}
	}
	if !task.UserID.Equal(userID) {
		return &domain.ErrPermissionDenied{}
	}
	task.DueDate = dueDate
	task.UpdatedAt = s.IClockManager.GetNow()
	if err := s.ITaskRepository.UpdateTask(ctx, task); err != nil {
		return &domain.ErrQueryFailed{}
	}
	return nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id string, userID string) error {
	if err := value.NewID(id).Validate(); err != nil {
		return err
	}
	if err := value.NewID(userID).Validate(); err != nil {
		return err
	}
	task, err := s.ITaskRepository.FindTaskByID(ctx, id)
	if err != nil {
		return &domain.ErrNotFound{Msg: "task not found"}
	}
	if !task.UserID.Equal(userID) {
		return &domain.ErrPermissionDenied{}
	}
	if err := s.ITaskRepository.DeleteTask(ctx, id); err != nil {
		return &domain.ErrQueryFailed{}
	}
	return nil
}
