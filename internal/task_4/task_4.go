package task_4

import (
	"context"
	"math/rand"
)

// RepeatFn repeatedly calls fn and sends results to the returned chan
func RepeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
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

// Take reads from `in` <= nums times, sends them to returned chan
func Take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
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
