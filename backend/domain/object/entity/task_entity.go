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

var ValidStatuses = []TaskStatus{StatusPending, StatusInProgress, StatusCompleted}

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

func NewTask(id *value.ID, userID *value.ID, title string, description *string, status TaskStatus, dueDate *time.Time) *Task {
	return &Task{
		ID:          id,
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      status,
		DueDate:     dueDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
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
	if !IsValidStatus(t.Status) {
		return &domain.ErrValidationFailed{Msg: "[ERROR] invalid task status value"}
	}
	return nil
}

func (t *Task) SetTitle(title string) {
	t.Title = title
	t.UpdatedAt = time.Now()
}

func (t *Task) SetDescription(description *string) {
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *Task) ChangeStatus(status TaskStatus) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

func IsValidStatus(s TaskStatus) bool {
	for _, validStatus := range ValidStatuses {
		if s == validStatus {
			return true
		}
	}
	return false
}
