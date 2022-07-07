package expire

import "time"

// Deadline is a type that will expire once the deadline has passed.
type Deadline[T any] struct {
	value    *T
	deadline time.Time
}

// Unwrap returns the underlying value or nil if the Deadline has passed.
func (t *Deadline[T]) Unwrap() *T {
	if time.Now().After(t.deadline) {
		return nil
	}

	return t.value
}

// WithDeadline returns a new expiring value based on a ttl.
func WithDeadline[T any](value T, deadline time.Time) Deadline[T] {
	e := Deadline[T]{
		value:    &value,
		deadline: deadline,
	}

	return e
}
