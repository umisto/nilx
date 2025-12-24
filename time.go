package nilx

import (
	"database/sql"
	"time"
)

func Time(v time.Time) sql.NullTime {
	return sql.NullTime{Time: v, Valid: true}
}

func TimePtr(v *time.Time) sql.NullTime {
	if v == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *v, Valid: true}
}
