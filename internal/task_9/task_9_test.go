package task_9

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_orDone_Chan_closed(t *testing.T) {
	ch := make(chan interface{})
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	var res []interface{}
	for v := range OrDone(context.Background(), ch) {
		res = append(res, v)
	}
	assert.Equal(t, []interface{}{0, 1, 2}, res)
}

func Test_orDone_Timeout(t *testing.T) {
	ch := make(chan interface{})
	defer close(ch)

	ctx, closeFn := context.WithTimeout(context.Background(), time.Second)
	defer closeFn()
	sendTil := 50000000
	go func() {

		for i := 0; i < sendTil; i++ {
			select {
			case <-ctx.Done():
				time.Sleep(2 * time.Second)
				return
			case ch <- i:
			}
		}
	}()

	var res []interface{}
	for v := range OrDone(ctx, ch) {
		res = append(res, v)
	}
	assert.True(t, len(res) >= 1)
	assert.True(t, len(res) < sendTil)
}
