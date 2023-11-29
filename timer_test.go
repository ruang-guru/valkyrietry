package valkyrietry

import (
	"testing"
	"time"
)

func TestTimerStart(t *testing.T) {
	timer := NewTimer()
	duration := 100 * time.Millisecond

	start := time.Now()
	timer.Start(duration)
	<-timer.C()

	if time.Since(start) < duration {
		t.Errorf("Timer fired before the expected duration")
	}
}

func TestTimerReset(t *testing.T) {
	timer := NewTimer()
	firstDuration := 50 * time.Millisecond
	secondDuration := 100 * time.Millisecond

	timer.Start(firstDuration)
	time.Sleep(30 * time.Millisecond)
	timer.Start(secondDuration)

	start := time.Now()
	<-timer.C()

	if elapsed := time.Since(start); elapsed < secondDuration {
		t.Errorf("Timer fired before the expected reset duration, elapsed: %v", elapsed)
	}
}

func TestTimerStop(t *testing.T) {
	timer := NewTimer()
	duration := 100 * time.Millisecond

	timer.Start(duration)
	timer.Stop()

	select {
	case <-timer.C():
		t.Errorf("Timer channel should not receive after being stopped")
	case <-time.After(150 * time.Millisecond):
	}
}
