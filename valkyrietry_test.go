package valkyrietry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestBasicRetrySuccess(t *testing.T) {
	ctx := context.Background()
	failureCount := 0
	maxFailures := 3
	retryFunc := func() error {
		if failureCount < maxFailures {
			failureCount++
			return errors.New("temporary error")
		}
		return nil
	}

	err := Do(ctx, retryFunc, WithMaxRetryAttempts(5))

	if err != nil {
		t.Errorf("Expected function to succeed, but it failed: %v", err)
	}
	if failureCount != maxFailures {
		t.Errorf("Expected %d failures, got %d", maxFailures, failureCount)
	}
}

func TestMaxRetryAttemptsExceeded(t *testing.T) {
	ctx := context.Background()
	retryFunc := func() error {
		return errors.New("permanent error")
	}

	err := Do(
		ctx,
		retryFunc,
		WithMaxRetryAttempts(2),
	)

	if err == nil || !errors.Is(err, ErrMaxRetryAttemptsExceeded) {
		t.Errorf("Expected ErrMaxRetryAttemptsExceeded, got %v", err)
	}
}

func TestRetryWithContextCancellation(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	retryFunc := func() error {
		time.Sleep(500 * time.Millisecond)
		return errors.New("error after delay")
	}

	err := Do(
		ctx,
		retryFunc,
	)

	if err == nil || err != context.DeadlineExceeded {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}
}

func TestJitterAndBackoff(t *testing.T) {
	ctx := context.Background()
	var retryDurations []time.Duration
	var lastRetryTime time.Time

	retryFunc := func() error {
		if !lastRetryTime.IsZero() {
			retryDurations = append(retryDurations, time.Since(lastRetryTime))
		}
		lastRetryTime = time.Now()
		return errors.New("retry error")
	}

	initialDelay := 500 * time.Millisecond
	maxAttempts := uint(3)
	backoffMultiplier := float32(1.5)
	jitterPercentage := float32(0.5) // 50% jitter

	_ = Do(
		ctx,
		retryFunc,
		WithMaxRetryAttempts(maxAttempts),
		WithRetryDelay(initialDelay),
		WithRetryBackoffMultiplier(backoffMultiplier),
		WithJitter(jitterPercentage),
	)

	baseDuration := initialDelay

	for i, duration := range retryDurations {
		jitter := time.Duration(float64(baseDuration) * float64(jitterPercentage))
		minDuration := baseDuration - jitter
		maxDuration := baseDuration + jitter

		if duration < minDuration || duration > maxDuration {
			t.Errorf("Retry interval %d is out of expected range: got %v, want between %v and %v", i, duration, minDuration, maxDuration)
		}

		baseDuration = time.Duration(float32(baseDuration) * backoffMultiplier)
	}
}

func TestRetryIntervalProgression(t *testing.T) {
	ctx := context.Background()
	var retryDurations []time.Duration
	var lastRetryTime time.Time

	retryFunc := func() error {
		if !lastRetryTime.IsZero() {
			retryDurations = append(retryDurations, time.Since(lastRetryTime))
		}
		lastRetryTime = time.Now()
		return errors.New("retry error")
	}

	initialDelay := 500 * time.Millisecond
	maxAttempts := uint(3)
	backoffMultiplier := float32(1.5)
	jitterPercentage := float32(0.5) // 50% jitter

	_ = Do(
		ctx,
		retryFunc,
		WithMaxRetryAttempts(maxAttempts),
		WithRetryDelay(initialDelay),
		WithRetryBackoffMultiplier(backoffMultiplier),
		WithJitter(jitterPercentage),
	)

	baseDuration := initialDelay

	for i, duration := range retryDurations {
		jitter := time.Duration(float64(baseDuration) * float64(jitterPercentage))
		minDuration := baseDuration - jitter
		maxDuration := baseDuration + jitter

		if duration < minDuration || duration > maxDuration {
			t.Errorf("Retry interval %d is out of expected range: got %v, want between %v and %v", i, duration, minDuration, maxDuration)
		}

		baseDuration = time.Duration(float32(baseDuration) * backoffMultiplier)
	}
}
