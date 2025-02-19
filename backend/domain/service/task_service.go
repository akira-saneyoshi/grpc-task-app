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
	taskStatus := entity.StatusPending
	if status != "" {
		taskStatus = entity.TaskStatus(status)
	}

	task := entity.NewTask(
		value.NewID(s.IIDManager.GenerateID()),
		value.NewID(userID),
		title,
		description,
		taskStatus,
		dueDate,
	)
	task.CreatedAt = now
	task.UpdatedAt = now

	if err := task.Validate(); err != nil {
		return "", err
	}
	createdID, err := s.ITaskRepository.CreateTask(ctx, task)
	if err != nil {
		return "", &domain.ErrQueryFailed{}
	}
	return createdID, nil
}

func (s *TaskService) UpdateTaskDetails(ctx context.Context, id string, userID string, title string, description *string, status string, dueDate *time.Time) error {
	// ID のバリデーション
	if err := value.NewID(id).Validate(); err != nil {
		return err
	}
	if err := value.NewID(userID).Validate(); err != nil {
		return err
	}

	// タスクの取得
	task, err := s.ITaskRepository.FindTaskByID(ctx, id)
	if err != nil {
		return &domain.ErrNotFound{Msg: "task not found"}
	}

	// 権限チェック (UserID が一致するか)
	if !task.UserID.Equal(userID) {
		return &domain.ErrPermissionDenied{}
	}

	// タスクの各フィールドを更新
	task.SetTitle(title)             // Title の更新
	task.SetDescription(description) // Description の更新

	// statusが空文字でない場合のみStatusを更新
	if status != "" {
		taskStatus := entity.TaskStatus(status)
		task.ChangeStatus(taskStatus) // Status の更新
	}

	// DueDate は nil でない場合のみ更新
	if dueDate != nil {
		task.DueDate = dueDate // DueDateの更新
	}

	task.UpdatedAt = s.IClockManager.GetNow() // UpdatedAt の更新

	// 更新されたタスクのバリデーション
	if err := task.Validate(); err != nil {
		return err
	}

	// データベースの更新
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
