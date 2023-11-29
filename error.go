package valkyrietry

import "fmt"

var (
	ErrMaxRetryAttemptsExceeded = fmt.Errorf("function is failed to retry after max attemps retries")
)
