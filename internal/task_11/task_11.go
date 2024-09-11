package task_11

import (
	"context"
	"github.com/konorlevich/course_goroutines_chanels/internal/task_9"
)

// bridge reads channels from ins and puts all the data to returned chanel
// till the end of data in channels or context cancellation
func bridge(ctx context.Context, ins <-chan <-chan interface{}) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)

		for {
			select {
			case <-ctx.Done():
				return
			case vCh, ok := <-ins:
				if !ok {
					return
				}
				for v := range task_9.OrDone(ctx, vCh) {
					ch <- v
				}
			}
		}
	}()

	return ch
}
