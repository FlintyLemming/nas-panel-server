package model

import (
	"time"
)

type Metric struct {
	ID        uint `gorm:"primaryKey"`
	ServerID  uint
	MetricID  uint
	Value     float64
	Timestamp time.Time
}
