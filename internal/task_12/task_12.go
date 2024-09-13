package task_12

import (
	"context"
	"errors"
	"sync"
)

// waiter runs all the f(ctx) in parallel.
// Number of parallel task to execute set via maxParallel during waiter initialization in newGroupWait
type waiter interface {
	// wait returns error if there was any in f(ctx) runs.
	// If there are more than one error, they wrapped with errors.Join
	wait() error
	run(ctx context.Context, f func(ctx context.Context) error)
}

type waitGroup struct {
	running chan struct{}

	errMu sync.Mutex
	err   error

	wg sync.WaitGroup
}

func (g *waitGroup) wait() error {
	g.wg.Wait()
	close(g.running)
	return g.err
}

func (g *waitGroup) run(ctx context.Context, fn func(ctx context.Context) error) {
	g.wg.Add(1)
	go func() {
		g.running <- struct{}{}
		go func() {
			defer func(running chan struct{}) { <-running; g.wg.Done() }(g.running)
			err := fn(ctx)
			if err == nil {
				return
			}
			g.errMu.Lock()
			defer g.errMu.Unlock()
			if g.err == nil {
				g.err = err
			} else {
				g.err = errors.Join(g.err, err)
			}
		}()
	}()
}

func newGroupWait(maxParallel int) waiter {
	return &waitGroup{running: make(chan struct{}, maxParallel)}
}
