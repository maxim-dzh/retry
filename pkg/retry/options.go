package retry

// Option ...
type Option func(r *retry)

// Attempts ...
func Attempts(attempts int) Option {
	return func(r *retry) {
		r.attempts = attempts
	}
}
