package valkyrietry

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

// RetryFunc
// Implement this function for retryable actions within the `Do` or `DoWithContext` functions.
// It is necessary to define this function if you want to utilize the
// retry mechanism inside `Do` or `DoWithContext`.
type RetryFunc func() error

// RetryFuncWithData
// Implement this function for retryable actions within the `DoWithData` or `DoWithDataAndContext` functions.
// It is necessary to define this function if you intend to use
// the retry mechanism inside `DoWithData` or `DoWithDataAndContext`.
type RetryFuncWithData[T any] func() (T, error)

type Valkyrietry struct {
	Configuration *Configuration

	ctx context.Context
}

func defaultValkyrietry(options ...Option) *Valkyrietry {
	defaultConfiguration := &Configuration{
		MaxRetryAttempts:       DefaultMaxRetryAttempt,
		InitialRetryDelay:      DefaultRetryDelay,
		RetryBackoffMultiplier: DefaultRetryBackoffMultiplier,
		JitterPercentage:       DefaultJitter,
	}

	defaultValue := &Valkyrietry{
		Configuration: defaultConfiguration,
		ctx:           context.Background(),
	}

	for _, opt := range options {
		opt(defaultConfiguration)
	}

	return defaultValue
}

func defaultValkyrietryWithContext(ctx context.Context, options ...Option) *Valkyrietry {
	defaultValue := defaultValkyrietry(options...)
	defaultValue.ctx = ctx

	return defaultValue
}

// Do
// Start the retry mechanism using the given context and continue running the process until the `RetryFunc`
// returns successfully (without error) or until the maximum number of retry attempts is exceeded.
//
// This function ensures that the given `RetryFunc` will run at least once.
func Do(ctx context.Context, f RetryFunc, options ...Option) error {
	valkyrietry := defaultValkyrietryWithContext(ctx, options...)

	currentAttempt := 0
	currentInterval := valkyrietry.Configuration.InitialRetryDelay

	// Initialize the timer to a zero value for
	// the first initialization.
	timer := time.NewTimer(0)

	defer func() {
		timer.Stop()
	}()

	var err error

	for {
		if currentAttempt > int(valkyrietry.Configuration.MaxRetryAttempts) {
			return errors.Join(err, ErrMaxRetryAttemptsExceeded)
		}

		err = f()

		if err == nil {
			return nil
		}

		retryInterval := valkyrietry.getRetryIntervalValue(currentInterval)
		currentInterval = time.Duration(float32(currentInterval) * valkyrietry.Configuration.RetryBackoffMultiplier)

		currentAttempt++

		// Reset the duration to match the retry interval duration.
		// Thus, we will adjust the timer interval for each retry.
		timer.Reset(retryInterval)

		select {
		case <-valkyrietry.ctx.Done():
			return valkyrietry.ctx.Err()
		case <-timer.C:
		}
	}
}

// DoWithData
// Start the retry mechanism with any given data to receive and a context, and continue running the process until the `RetryFunc`
// successfully returns with the data (without error) or until the maximum number of retry attempts is exceeded.
//
// This function ensures that the given `RetryFunc` will run at least once.
func DoWithData[T any](ctx context.Context, f RetryFuncWithData[T], options ...Option) (T, error) {
	valkyrietry := defaultValkyrietryWithContext(ctx, options...)

	currentAttempt := 0
	currentInterval := valkyrietry.Configuration.InitialRetryDelay

	// Initialize the timer to a zero value for
	// the first initialization.
	timer := time.NewTimer(0)

	defer func() {
		timer.Stop()
	}()

	var response T
	var err error

	for {
		if currentAttempt > int(valkyrietry.Configuration.MaxRetryAttempts) {
			return response, errors.Join(err, ErrMaxRetryAttemptsExceeded)
		}

		response, err = f()

		if err == nil {
			return response, nil
		}

		retryInterval := valkyrietry.getRetryIntervalValue(currentInterval)
		currentInterval = time.Duration(float32(currentInterval) * valkyrietry.Configuration.RetryBackoffMultiplier)

		currentAttempt++

		// Reset the duration to match the retry interval duration.
		// Thus, we will adjust the timer interval for each retry.
		timer.Reset(retryInterval)

		select {
		case <-valkyrietry.ctx.Done():
			return response, valkyrietry.ctx.Err()
		case <-timer.C:
		}
	}
}

func (v *Valkyrietry) getRetryIntervalValue(currentInterval time.Duration) time.Duration {
	jitterInterval := v.Configuration.JitterPercentage * float32(currentInterval)

	maxRetryInterval := float32(currentInterval) + jitterInterval
	minRetryInterval := float32(currentInterval) - jitterInterval

	randomValue := rand.Float32()

	return time.Duration(minRetryInterval + (randomValue * (maxRetryInterval - minRetryInterval + 1)))
}
