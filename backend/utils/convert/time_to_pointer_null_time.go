package convert

import (
	"database/sql"
	"time"
)

// timeToPointerNullTime is a function that converts *time.Time to sql.NullTime.
func TimeToPointerNullTime(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{Time: *t, Valid: true}
	}
	return sql.NullTime{}
}
