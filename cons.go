package valkyrietry

import "time"

const (
	DefaultMaxRetryAttempt        = 5
	DefaultRetryDelay             = time.Duration(0.5 * float64(time.Second))
	DefaultRetryBackoffMultiplier = 1.5
	DefaultJitter                 = 0.5
)
