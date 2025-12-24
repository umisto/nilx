package nilx

import "database/sql"

func Int(v int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(v), Valid: true}
}

func Int64(v int64) sql.NullInt64 {
	return sql.NullInt64{Int64: v, Valid: true}
}

func Int32(v int32) sql.NullInt32 {
	return sql.NullInt32{Int32: v, Valid: true}
}

func IntPtr(v *int) sql.NullInt64 {
	if v == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: int64(*v), Valid: true}
}

func Int64Ptr(v *int64) sql.NullInt64 {
	if v == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: *v, Valid: true}
}

func Int32Ptr(v *int32) sql.NullInt32 {
	if v == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *v, Valid: true}
}
