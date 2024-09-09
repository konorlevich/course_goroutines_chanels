package task_3

import "context"

// generator
// receives a context and a slice if numbers that put into the channel returned
func generator(ctx context.Context, numbers ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range numbers {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()

	return out
}

// squarer
// receives a context and a channel
//
// reads all the numbers from the nums chan, square them and return to output chan
func squarer(ctx context.Context, nums <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n * n:
			}
		}
	}()

	return out
}
