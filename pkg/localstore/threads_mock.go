package localstore

import (
	"testing"

	"context"

	sourcegraph "sourcegraph.com/sourcegraph/sourcegraph/pkg/api"
)

type MockThreads struct {
	Get    func(ctx context.Context, id int64) (*sourcegraph.Thread, error)
	Create func(ctx context.Context, newThread *sourcegraph.Thread) (*sourcegraph.Thread, error)
}

func (s *MockThreads) MockGet_Return(t *testing.T, returns *sourcegraph.Thread, returnsErr error) (called *bool) {
	called = new(bool)
	s.Get = func(ctx context.Context, id int64) (*sourcegraph.Thread, error) {
		*called = true
		return returns, returnsErr
	}
	return called
}

func (s *MockThreads) MockCreate_Return(t *testing.T, returns *sourcegraph.Thread, returnsErr error) (called *bool, calledWith *sourcegraph.Thread) {
	called, calledWith = new(bool), &sourcegraph.Thread{}
	s.Create = func(ctx context.Context, newThread *sourcegraph.Thread) (*sourcegraph.Thread, error) {
		*called = true
		return returns, returnsErr
	}
	return called, calledWith
}
