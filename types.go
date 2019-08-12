package common

import (
	"fmt"
	"time"
)

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

// DateTime .
type DateTime time.Time

// MarshalJSON .
func (t DateTime) MarshalJSON() ([]byte, error) {
	t1 := time.Time(t)
	if t1.IsZero() {
		return []byte("null"), nil
	}
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
