package convert

import (
	"database/sql"
	"time"
)

// convertNullTime is a helper function that converts sql.NullTime to time.Time.
func ConvertNullTime(nullTime sql.NullTime) time.Time {
	if nullTime.Valid {
		return nullTime.Time
	}
	return time.Time{}
}
