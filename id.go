package common

// ID .
type ID int64

// Int64 .
func (id ID) Int64() int64 {
	return int64(id)
}

// Str .
func (id ID) Str() string {
	return Int64ToStr(int64(id))
}

// MarshalText .
func (id ID) MarshalText() ([]byte, error) {
	return []byte(id.Str()), nil
}

// Int64ToID .
func Int64ToID(i int64) ID {
	return ID(i)
}

// StrToID .
func StrToID(s string) ID {
	return Int64ToID(StrToInt64(s, 0))
}
