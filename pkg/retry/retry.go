package retry

import "context"

const (
	defaultAttempts = 3
)

type returnErrFunc func() (err error)

// Retry ...
type Retry interface {
	Do(ctx context.Context, f returnErrFunc) (err error)
}

type retry struct {
	attempts int
}

// Do ...
func (r *retry) Do(ctx context.Context, f returnErrFunc) (err error) {
	if err = ctx.Err(); err != nil {
		return
	}
	var n int
	for n < r.attempts {
		err = f()
		if err != nil {
			// TODO: implement
		}
	}
	return
}

// NewRetry ...
func NewRetry(opts ...Option) Retry {
	r := retry{
		attempts: defaultAttempts,
	}
	for _, opt := range opts {
		opt(&r)
	}
	return &r
}
