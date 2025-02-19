package entity

import (
	"time"

	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted  TaskStatus = "completed"
)

type Task struct {
	ID          *value.ID
	UserID      *value.ID
	Title       string
	Description *string
	Status      TaskStatus
	DueDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) Validate() error {
	if err := t.ID.Validate(); err != nil {
		return err
	}
	if err := t.UserID.Validate(); err != nil {
		return err
	}
	if t.Title == "" {
		return &domain.ErrValidationFailed{Msg: "[ERROR] title is empty"}
	}
	if t.Status != StatusPending && t.Status != StatusInProgress && t.Status != StatusCompleted {
		return &domain.ErrValidationFailed{Msg: "[ERROR] invalid status value"}
	}
	return nil
}
