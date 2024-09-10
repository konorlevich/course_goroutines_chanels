package task_10

import (
	"context"
	"github.com/konorlevich/course_goroutines_chanels/internal/task_4"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_tee_from_task(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	i := 0
	inc := func() interface{} {
		i++
		return i
	}

	out1, out2 := tee(ctx, task_4.Take(ctx, task_4.RepeatFn(ctx, inc), 3))
	var res1, res2 []interface{}
	for val1 := range out1 {
		res1 = append(res1, val1)
		res2 = append(res2, <-out2)
	}
	exp := []interface{}{1, 2, 3}
	assert.Equal(t, exp, res1)
	assert.Equal(t, exp, res2)
}

func Test_tee_Timeout(t *testing.T) {
	in := make(chan interface{})
	defer close(in)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sendTil := 500000
	go func() {
		for i := 0; i < sendTil; i++ {
			select {
			case <-ctx.Done():
				time.Sleep(5 * time.Second)
				return
			case in <- i:
				time.Sleep(time.Second)
			}
		}
	}()
	out1, out2 := tee(ctx, in)
	res1, res2 := make([]interface{}, 0, 200), make([]interface{}, 0, 200)

	for i1 := range out1 {
		res1 = append(res1, i1)
		res2 = append(res2, <-out2)
	}
	_, ok := <-out1
	assert.False(t, ok)
	_, ok = <-out2
	assert.False(t, ok)

	assert.Equal(t, res1, res2)
	assert.Less(t, len(res1), sendTil)
}
