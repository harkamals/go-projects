package main

import (
	"time"
)

// type TimeInstant time.Time

func (t *TimeInstant) UnmarshalText() string {
	return "Hello"
}
