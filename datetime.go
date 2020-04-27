package common

import (
	"fmt"
	"time"
)

// DateTime .
type DateTime time.Time

// IsZero .
func (t DateTime) IsZero() bool {
	t1 := time.Time(t)
	return t1.IsZero()
}

// MarshalJSON .
func (t DateTime) MarshalJSON() ([]byte, error) {
	t1 := time.Time(t)
	if t1.IsZero() {
		return []byte("null"), nil
	}
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// Now 当前时间
func Now() DateTime {
	return DateTime(time.Now())
}
