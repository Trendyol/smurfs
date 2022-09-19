package process

import "context"

type Exec interface {
	// Run executes the command with the given arguments
	Run(ctx context.Context, command string, args ...string) error
}

type exec struct{}

func New() Exec {
	return &exec{}
}

func (e *exec) Run(ctx context.Context, command string, args ...string) error {
	panic("implement me")
}
