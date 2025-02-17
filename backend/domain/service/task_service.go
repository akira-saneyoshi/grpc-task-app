package service

import (
	"context"
	"time"

	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	"github.com/akira-saneyoshi/task-app/domain/repository"
	"github.com/akira-saneyoshi/task-app/utils/clock"
	"github.com/akira-saneyoshi/task-app/utils/identification"
)

type ITaskService interface {
	FindTaskByID(ctx context.Context, id string) (*entity.Task, error)
	FindTasksByUserID(ctx context.Context, userID string) ([]*entity.Task, error)
	CreateTask(ctx context.Context, userID string, title string, description *string, status string, dueDate *time.Time) (string, error)
	UpdateTaskDetails(ctx context.Context, id string, userID string, title string, description *string, status string, dueDate *time.Time) error
	UpdateTaskStatus(ctx context.Context, id string, userID string, status string) error
	UpdateTaskDueDate(ctx context.Context, id string, userID string, dueDate *time.Time) error
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
		return nil, &domain.ErrNotFound{Msg: "[ERROR] task not found"}
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
	taskStatus := entity.StatusPending
	if status != "" {
		taskStatus = entity.Status(status)
	}

	arg := &entity.Task{
		ID:          value.NewID(s.IIDManager.GenerateID()),
		UserID:      value.NewID(userID),
		Title:       title,
		Description: description,
		Status:      taskStatus,
		DueDate:     time.Time{},
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if dueDate != nil {
		arg.DueDate = *dueDate
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

func (s *TaskService) UpdateTaskDetails(ctx context.Context, id string, userID string, title string, description *string, status string, dueDate *time.Time) error {
	if err := value.NewID(id).Validate(); err != nil {
		return err
	}
	if err := value.NewID(userID).Validate(); err != nil {
		return err
	}
	task, err := s.ITaskRepository.FindTaskByID(ctx, id)
	if err != nil {
		return &domain.ErrNotFound{Msg: "[ERROR] task not found"}
	}
	if !task.UserID.Equal(userID) {
		return &domain.ErrPermissionDenied{}
	}
	task.Title = title
	task.Description = description
	taskStatus := entity.Status(status)
	if status != "" && !entity.IsValidStatus(taskStatus) {
		return &domain.ErrValidationFailed{Msg: "[ERROR] invalid status value"}
	}
	if status != "" {
		task.Status = taskStatus
	}
	task.DueDate = time.Time{}
	if dueDate != nil {
		task.DueDate = *dueDate
	}
	task.UpdatedAt = s.IClockManager.GetNow()
	if err := task.Validate(); err != nil {
		return err
	}
	if err := s.ITaskRepository.UpdateTask(ctx, task); err != nil {
		return &domain.ErrQueryFailed{}
	}
	return nil
}

func (s *TaskService) UpdateTaskStatus(ctx context.Context, id string, userID string, status string) error {
	if err := value.NewID(id).Validate(); err != nil {
		return err
	}
	if err := value.NewID(userID).Validate(); err != nil {
		return err
	}
	task, err := s.ITaskRepository.FindTaskByID(ctx, id)
	if err != nil {
		return &domain.ErrNotFound{Msg: "[ERROR] task not found"}
	}
	if !task.UserID.Equal(userID) {
		return &domain.ErrPermissionDenied{}
	}
	taskStatus := entity.Status(status)
	if !entity.IsValidStatus(taskStatus) {
		return &domain.ErrValidationFailed{Msg: "[ERROR] invalid status value"}
	}
	task.Status = taskStatus
	task.UpdatedAt = s.IClockManager.GetNow()
	if err := s.ITaskRepository.UpdateTask(ctx, task); err != nil {
		return &domain.ErrQueryFailed{}
	}
	return nil
}

func (s *TaskService) UpdateTaskDueDate(ctx context.Context, id string, userID string, dueDate *time.Time) error {
	if err := value.NewID(id).Validate(); err != nil {
		return err
	}
	if err := value.NewID(userID).Validate(); err != nil {
		return err
	}
	task, err := s.ITaskRepository.FindTaskByID(ctx, id)
	if err != nil {
		return &domain.ErrNotFound{Msg: "[ERROR] task not found"}
	}
	if !task.UserID.Equal(userID) {
		return &domain.ErrPermissionDenied{}
	}
	task.DueDate = time.Time{}
	if dueDate != nil {
		task.DueDate = *dueDate
	}
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
		return &domain.ErrNotFound{Msg: "[ERROR] task not found"}
	}
	if !task.UserID.Equal(userID) {
		return &domain.ErrPermissionDenied{}
	}
	if err := s.ITaskRepository.DeleteTask(ctx, id); err != nil {
		return &domain.ErrQueryFailed{}
	}
	return nil
}
