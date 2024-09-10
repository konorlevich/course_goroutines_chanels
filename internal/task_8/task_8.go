package task_8

import (
	"context"
	"math/rand"
	"time"
)

var timeout = 100 * time.Millisecond

// executeTaskWithTimeout
//
// receives a context, runs executeTask
//
// finishes on executeTask end or context cancellation
//
// returns a context error if any
func executeTaskWithTimeout(ctx context.Context) error {
	ch := make(chan struct{})
	go func() {
		executeTask(timeout)
		ch <- struct{}{}
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ch:
	}

	return nil
}

// executeTask can hang for a while
func executeTask(timeout time.Duration) {
	time.Sleep(time.Duration(rand.Intn(3)) * timeout)
}
