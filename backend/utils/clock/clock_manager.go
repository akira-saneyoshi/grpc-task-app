package clock

import "time"

type IClockManager interface {
	GetNow() time.Time
}

type ClockManager struct{}

func NewClockManager() *ClockManager {
	return &ClockManager{}
}

func (m *ClockManager) GetNow() time.Time {
	return time.Now().UTC()
}
