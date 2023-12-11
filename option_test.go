package valkyrietry

import (
	"testing"
	"time"
)

func TestWithMaxRetryAttempts(t *testing.T) {
	expectedAttempts := uint(5)
	config := &Configuration{}
	option := WithMaxRetryAttempts(expectedAttempts)
	option(config)

	if config.MaxRetryAttempts != expectedAttempts {
		t.Errorf("WithMaxRetryAttempts() = %v, want %v", config.MaxRetryAttempts, expectedAttempts)
	}
}

func TestWithRetryDelay(t *testing.T) {
	expectedDelay := 100 * time.Millisecond
	config := &Configuration{}
	option := WithRetryDelay(expectedDelay)
	option(config)

	if config.InitialRetryDelay != expectedDelay {
		t.Errorf("WithRetryDelay() = %v, want %v", config.InitialRetryDelay, expectedDelay)
	}
}

func TestWithRetryBackoffMultiplier(t *testing.T) {
	expectedMultiplier := float32(2.0)
	config := &Configuration{}
	option := WithRetryBackoffMultiplier(expectedMultiplier)
	option(config)

	if config.RetryBackoffMultiplier != expectedMultiplier {
		t.Errorf("WithRetryBackoffMultiplier() = %v, want %v", config.RetryBackoffMultiplier, expectedMultiplier)
	}
}

func TestWithJitter(t *testing.T) {
	expectedJitter := float32(0.25)
	config := &Configuration{}
	option := WithJitter(expectedJitter)
	option(config)

	if config.JitterPercentage != expectedJitter {
		t.Errorf("WithJitter() = %v, want %v", config.JitterPercentage, expectedJitter)
	}
}
