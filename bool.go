package nilx

import "database/sql"

func Bool(v bool) sql.NullBool {
	return sql.NullBool{Bool: v, Valid: true}
}

func BoolPtr(v *bool) sql.NullBool {
	if v == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{Bool: *v, Valid: true}
}
