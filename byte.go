package nilx

import "database/sql"

func Byte(v *byte) sql.NullByte {
	if v == nil {
		return sql.NullByte{}
	}
	return sql.NullByte{Byte: *v, Valid: true}
}
