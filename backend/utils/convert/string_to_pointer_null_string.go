package convert

import (
	"database/sql"
)

// stringToPointerNullString is a function that converts *string to sql.NullString.
func StringToPointerNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{}
}
