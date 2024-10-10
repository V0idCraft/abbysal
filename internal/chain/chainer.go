package chain

import "context"

type Executor interface {
	// Responsibility of chain pattern
	Execute(ctx context.Context) error
	SetNext(Executor)
}
