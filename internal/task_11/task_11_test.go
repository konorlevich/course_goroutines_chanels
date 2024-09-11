package task_11

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_bridge_From_task(t *testing.T) {
	genVals := func() <-chan <-chan interface{} {
		out := make(chan (<-chan interface{}))
		go func() {
			defer close(out)
			for i := 0; i < 3; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				out <- stream
			}
		}()
		return out
	}

	var res []interface{}
	for v := range bridge(context.Background(), genVals()) {
		res = append(res, v)
	}

	assert.Equal(t, []interface{}{0, 1, 2}, res)
}

func Test_bridge_Cancelled_ctx(t *testing.T) {
	genVals := func() <-chan <-chan interface{} {
		out := make(chan (<-chan interface{}))
		go func() {
			defer close(out)
			for i := 0; i < 3; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				out <- stream
			}
		}()
		return out
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	vCh := bridge(ctx, genVals())
	v, ok := <-vCh
	assert.True(t, ok)
	assert.Equal(t, 0, v)
	v, ok = <-vCh
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	cancel()
	v, ok = <-vCh
	assert.False(t, ok)
}
