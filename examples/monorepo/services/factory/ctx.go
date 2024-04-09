package factory

import (
	"context"
	"time"
)

var _ context.Context = (*sf)(nil)

// Deadline implements context.Context.
func (s *sf) Deadline() (deadline time.Time, ok bool) {
	return s.Context().Deadline()
}

// Done implements context.Context.
func (s *sf) Done() <-chan struct{} {
	return s.Context().Done()
}

// Err implements context.Context.
func (s *sf) Err() error {
	return s.Context().Err()
}

// Value implements context.Context.
func (s *sf) Value(key any) any {
	return s.Context().Value(key)
}
