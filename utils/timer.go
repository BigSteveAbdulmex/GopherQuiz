package utils

import (
	"context"
	"fmt"
	"time"
)

// These variables help us keep track of time
var startTime time.Time       // when the quiz starts
var endTime time.Time         // when the quiz ends
var elapsedTime time.Duration // total time taken

// StartTimer begins the quiz timer.
// This is useful to calculate how long a user takes to finish the quiz.
func StartTimer() {
	startTime = time.Now()
}

// EndTimer stops the quiz timer.
// It saves how much time has passed since the timer started.
func EndTimer() {
	endTime = time.Now()
	elapsedTime = endTime.Sub(startTime)
}

// GetElapsedTime returns how many seconds the user spent on the quiz.
// We use this to show time taken at the end.
func GetElapsedTime() int {
	return int(elapsedTime.Seconds())
}

// StartCountdownContext starts a countdown using a context.
// After the time runs out, the context is automatically cancelled.
// Useful for stopping the quiz when time is up.
func StartCountdownContext(durationSeconds int) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(durationSeconds)*time.Second)
	return ctx, cancel
}

// FormatDuration converts seconds into a nice-looking format like 00:02:30 (hh:mm:ss).
// It helps us show the time limit or time taken in a way that's easy to read.
func FormatDuration(seconds int) string {
	duration := time.Duration(seconds) * time.Second
	return fmt.Sprintf("%02d:%02d:%02d", int(duration.Hours()), int(duration.Minutes())%60, int(duration.Seconds())%60)
}

// SecondsToDuration is a helper that converts seconds into time.Duration type.
// We use this when creating timers that need a duration.
func SecondsToDuration(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}
