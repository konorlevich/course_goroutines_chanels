package task_10

import (
	"context"
	"github.com/konorlevich/course_goroutines_chanels/internal/task_9"
)

// tee
// proxy data from the in chan to two returned channels (exact same data in both channels),
// while the `in` chan is open and context is not cancelled
//
// Hint: use orDone function from previous task for simplification
func tee(ctx context.Context, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1, out2 := make(chan interface{}), make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)
		for v := range task_9.OrDone(ctx, in) {
			select {
			case <-ctx.Done():
				return
			default:
				out1 <- v
				out2 <- v
			}
		}
	}()

	return out1, out2
}
