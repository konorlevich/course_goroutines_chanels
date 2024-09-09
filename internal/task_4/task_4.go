package task_4

import (
	"context"
	"math/rand"
)

// repeatFn repeatedly calls fn and sends results to the returned chan
func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case out <- fn():
			}
		}
	}()

	return out
}

// take reads from `in` <= nums times, sends them to returned chan
func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case data, ok := <-in:
				if !ok {
					return
				}
				out <- data
			}
		}
	}()
	return out
}

func fnRand() interface{} { return rand.Int() }
