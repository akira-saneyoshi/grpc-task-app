package convert

import (
	"database/sql"
)

// nullStringToPointer is a function that converts sql.NullString to *string.
func NullStringToPointer(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}
