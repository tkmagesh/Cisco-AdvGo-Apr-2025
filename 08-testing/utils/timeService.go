package utils

import "time"

type TimeService struct {
}

func (ts TimeService) GetCurrent() time.Time {
	return time.Now()
}
