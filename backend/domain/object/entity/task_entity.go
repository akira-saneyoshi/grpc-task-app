package entity

import (
	"time"

	"github.com/akira-saneyoshi/task-app/domain"
	"github.com/akira-saneyoshi/task-app/domain/object/value"
	task "github.com/akira-saneyoshi/task-app/domain/object/value/task"
)

type Task struct {
	ID          *value.ID   `json:"id"`
	UserID      *value.ID   `json:"user_id"`
	Title       *task.Title `json:"title"`
	Description *string     `json:"description"`
	Status      Status      `json:"status"`
	DueDate     *time.Time  `json:"due_date"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
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
	if err := t.Title.Validate(); err != nil {
		return err
	}
	if !IsValidStatus(t.Status) {
		return &domain.ErrValidationFailed{Msg: "invalid status value"}
	}
	return nil
}

func IsValidStatus(s Status) bool { // 関数名を IsValidStatus に変更 (先頭を大文字に)
	for _, validStatus := range ValidStatuses {
		if s == validStatus {
			return true
		}
	}
	return false
}
