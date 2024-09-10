package task_9

import (
	"context"
)

// orDone, которая направляет данные из канала in в возвращаемый канал,
// пока канал in открыт и контекст не отменен
func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		for i := range in {
			select {
			case <-ctx.Done():
				return
			default:
				out <- i
			}
		}
	}()

	return out
}
