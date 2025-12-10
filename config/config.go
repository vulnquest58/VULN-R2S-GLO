package config

import "time"

type Config struct {
	MaxConcurrent  int
	ScanHours      int
	WAFBypassKB    int
	RequestTimeout time.Duration
	PauseInterval  int // every N requests
	PauseDuration  time.Duration
}
