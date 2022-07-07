package expire

import "time"

// TTL is a type that will expire once the ttl has passed.
type TTL[T any] struct {
	value   *T
	created time.Time
	ttl     time.Duration
}

// Unwrap returns the underlying value or nil if the TTL has passed.
func (t *TTL[T]) Unwrap() *T {
	if t.created.Before(time.Now().Add(t.ttl * -1)) {
		return nil
	}

	return t.value
}

// WithTTL returns a new expiring value based on a ttl.
func WithTTL[T any](value T, ttl time.Duration) TTL[T] {
	e := TTL[T]{
		value:   &value,
		created: time.Now(),
		ttl:     ttl,
	}

	return e
}
