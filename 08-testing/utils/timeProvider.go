package utils

import "time"

type TimeProvider interface {
	GetCurrent() time.Time
}
