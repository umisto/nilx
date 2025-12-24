package nilx

import "github.com/google/uuid"

func UUID(v uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{UUID: v, Valid: true}
}

func UUIDPtr(v *uuid.UUID) uuid.NullUUID {
	if v == nil {
		return uuid.NullUUID{}
	}
	return UUID(*v)
}

func UUIDString(s string) (uuid.NullUUID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return uuid.NullUUID{}, err
	}
	return uuid.NullUUID{UUID: u, Valid: true}, nil
}

func UUIDStringPtr(s *string) (uuid.NullUUID, error) {
	if s == nil {
		return uuid.NullUUID{}, nil
	}
	return UUIDString(*s)
}
