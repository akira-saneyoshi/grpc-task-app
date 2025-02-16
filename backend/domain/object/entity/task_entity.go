package entity

import (
	"time"

	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
)

type Task struct {
	ID          *value.ID  `json:"id"`
	UserID      *value.ID  `json:"user_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      Status     `json:"status"`
	DueDate     *time.Time `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Status はタスクの状態を表すENUM型です。
type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
)

var ValidStatuses = []Status{StatusPending, StatusInProgress, StatusCompleted}

func (t *Task) Validate() error {
	if err := t.ID.Validate(); err != nil {
		return err
	}
	if err := t.UserID.Validate(); err != nil {
		return err
	}
	if t.Title == "" {
		return &domain.ErrValidationFailed{Msg: "title is empty"}
	}
	if !IsValidStatus(t.Status) {
		return &domain.ErrValidationFailed{Msg: "invalid status value"}
	}
	return nil
}

func IsValidStatus(s Status) bool {
	for _, validStatus := range ValidStatuses {
		if s == validStatus {
			return true
		}
	}
	return false
}
