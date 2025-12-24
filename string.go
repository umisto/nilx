package nilx

import (
	"database/sql"
)

func String(v string) sql.NullString {
	return sql.NullString{String: v, Valid: true}
}

func StringPtr(v *string) sql.NullString {
	if v == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *v, Valid: true}
}
