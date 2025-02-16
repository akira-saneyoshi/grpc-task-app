package convert

import (
	"database/sql"
	"time"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/model/db"
)

func ConvertNullString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
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

func NewSQLNullTime(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{
			Time:  *t,
			Valid: true,
		}
	} else {
		return sql.NullTime{
			Valid: false,
		}
	}
}
