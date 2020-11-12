package retry

import (
	"context"
	"time"
)

const (
	defaultAttempts = 3
	defaultDelay    = 5 * time.Second
)

type returnErrFunc func() (err error)

// Retry ...
type Retry interface {
	Do(ctx context.Context, f returnErrFunc) (err error)
}

type retry struct {
	attempts int
	delay    time.Duration
}

// Do ...
func (r *retry) Do(ctx context.Context, f returnErrFunc) (err error) {
	var (
		actualAttempts int
		ticker         = time.NewTicker(r.delay)
	)
	defer ticker.Stop()
	for actualAttempts < r.attempts {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err = f()
			if err == nil {
				return
			}

			actualAttempts++
			ticker.Reset(r.delay)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-ticker.C:
				continue
			}
		}
	}
	return
}

// NewRetry ...
func NewRetry(opts ...Option) Retry {
	r := retry{
		attempts: defaultAttempts,
		delay:    defaultDelay,
	}
	for _, opt := range opts {
		opt(&r)
	}
	return &r
}
