package convert

import (
	"database/sql"
	"time"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/model/db"
	task_v1 "github.com/akira-saneyoshi/task-app/interfaces/proto/task/v1"
)

func ConvertNullString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func ConvertNullTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

func ConvertNullTimePtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

func ConvertNullTimeValue(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

func ConvertNullStatus(ns db.NullTasksStatus) entity.Status {
	if ns.Valid {
		return entity.Status(ns.TasksStatus)
	}
	return entity.StatusPending
}

func NewSQLNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{
			String: *s,
			Valid:  true,
		}
	} else {
		return sql.NullString{
			Valid: false,
		}
	}
}

func NewSQLNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

func ConvertStatus(entityStatus entity.Status) task_v1.TaskStatus {
	switch entityStatus {
	case entity.StatusPending:
		return task_v1.TaskStatus_TASK_STATUS_PENDING_UNSPECIFIED
	case entity.StatusInProgress:
		return task_v1.TaskStatus_TASK_STATUS_IN_PROGRESS
	case entity.StatusCompleted:
		return task_v1.TaskStatus_TASK_STATUS_COMPLETED
	default:
		return task_v1.TaskStatus_TASK_STATUS_PENDING_UNSPECIFIED
	}
}

func ConvertTaskStatusToEntityStatus(protoStatus task_v1.TaskStatus) entity.Status {
	switch protoStatus {
	case task_v1.TaskStatus_TASK_STATUS_PENDING_UNSPECIFIED:
		return entity.StatusPending
	case task_v1.TaskStatus_TASK_STATUS_IN_PROGRESS:
		return entity.StatusInProgress
	case task_v1.TaskStatus_TASK_STATUS_COMPLETED:
		return entity.StatusCompleted
	default:
		return entity.StatusPending
	}
}

func ConvertEntityStatusToString(entityStatus entity.Status) string {
	switch entityStatus {
	case entity.StatusPending:
		return "pending"
	case entity.StatusInProgress:
		return "in_progress"
	case entity.StatusCompleted:
		return "completed"
	default:
		return "pending"
	}
}
