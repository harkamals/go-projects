package self_learn

import (
	"time"
)

// type TimeInstant time.Time

func (t *TimeInstant) UnmarshalText() string {
	return "Hello"
}
