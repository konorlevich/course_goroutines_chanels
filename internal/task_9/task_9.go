package task_9

import (
	"context"
)

// orDone proxy data from the in chan to returned chan,
// while the in chan is open and context is not cancelled
func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- v:
				case <-ctx.Done():
				}
			default:

			}
		}
	}()

	return out
}
