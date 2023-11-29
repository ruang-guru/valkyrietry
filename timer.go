package valkyrietry

import "time"

type Timer struct {
	timer *time.Timer
}

func NewTimer() *Timer {
	return &Timer{}
}

// Start
// Set the timer for the specified duration.
// If the current timer is nil, initialize a new one;
// otherwise, reset it to the new duration.
func (t *Timer) Start(duration time.Duration) {
	if t.timer == nil {
		t.timer = time.NewTimer(duration)
		return
	}

	t.timer.Reset(duration)
}

// Stop
// Stop the current timer.
func (t *Timer) Stop() {
	if t.timer != nil {
		t.timer.Stop()
	}
}

// C
// Retrieve the channel when either the timer stops or the timer completes.
func (t *Timer) C() <-chan time.Time {
	return t.timer.C
}
