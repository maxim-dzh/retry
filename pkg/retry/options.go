package retry

import "time"

// Option ...
type Option func(r *retry)

// Attempts ...
func Attempts(attempts int) Option {
	return func(r *retry) {
		r.attempts = attempts
	}
}

// Delay ...
func Delay(delay time.Duration) Option {
	return func(r *retry) {
		r.delay = delay
	}
}
