package convert

import (
	"database/sql"
	"time"
)

// nullTimeToTimePointer is a function that converts sql.NullTime to *time.Time.
func NullTimeToTimePointer(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
