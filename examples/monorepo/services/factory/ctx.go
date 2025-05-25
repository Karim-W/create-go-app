package factory

import (
	"context"
	"time"
)

var _ context.Context = (*sf)(nil)

func (s *sf) Deadline() (deadline time.Time, ok bool) {
	return s.Context().Deadline()
}

func (s *sf) Done() <-chan struct{} {
	return s.Context().Done()
}

func (s *sf) Err() error {
	return s.Context().Err()
}

func (s *sf) Value(key any) any {
	return s.Context().Value(key)
}

// Context() Returns the context of the request
func (s *sf) Context() context.Context {
	return s.ctx
}
