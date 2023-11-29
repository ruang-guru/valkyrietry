package valkyrietry

import (
	"time"
)

type Configuration struct {
	MaxRetryAttempts       uint
	InitialRetryDelay      time.Duration
	RetryBackoffMultiplier float32
	JitterPercentage       float32
}

// option is a function option used to configure a Valkyrietry.
type Option func(c *Configuration)

// WithMaxRetryAttempts
// Set the maximum number of retry attempts for the retry mechanism.
//
// if you set the attempt to 0, it means it will run until the process succed
func WithMaxRetryAttempts(attempt uint) Option {
	return func(c *Configuration) {
		c.MaxRetryAttempts = attempt
	}
}

// WithRetryDelay
// Set the initial duration value for the first retry.
func WithRetryDelay(delay time.Duration) Option {
	return func(c *Configuration) {
		c.InitialRetryDelay = delay
	}
}

// WithRetryBackoffMultiplier
// Set the multiplier for each failed retry attempt.
// Formula: initial retry delay * multiplier.
func WithRetryBackoffMultiplier(multiplier float32) Option {
	return func(c *Configuration) {
		c.RetryBackoffMultiplier = multiplier
	}
}

// WithJitter
// Set the percentage of jitter value to determine the lowest and highest
// random values.
func WithJitter(percentage float32) Option {
	return func(c *Configuration) {
		c.JitterPercentage = percentage
	}
}
