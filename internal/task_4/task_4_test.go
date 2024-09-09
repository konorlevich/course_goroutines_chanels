package task_4

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_From_The_task(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var res []interface{}
	for num := range take(ctx, repeatFn(ctx, fnRand), 3) {
		res = append(res, num)
	}
	assert.Equal(t, 3, len(res))
}

func Test_repeatFn(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := repeatFn(ctx, fnRand)
	for i := 0; i < 100; i++ {
		num, ok := <-ch
		assert.True(t, ok)
		_, ok = (num).(int)
		assert.True(t, ok)
	}
	cancel()
	_, ok := <-ch
	if ok { // if the last num stuck in the chan
		_, ok = <-ch
	}
	assert.False(t, ok)
}

func Test_take(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch1 := repeatFn(ctx, fnRand)

	ch2 := take(ctx, ch1, 100)
	for i := 0; i < 100; i++ {
		num, ok := <-ch2
		assert.True(t, ok)
		_, ok = (num).(int)
		assert.True(t, ok)
	}
	_, ok := <-ch2
	assert.False(t, ok)

	_, ok = <-ch1
	assert.True(t, ok)
	cancel()
	_, ok = <-ch1 //take the last sent val
	_, ok = <-ch1
	assert.False(t, ok)
}

func Test_take_with_different_ctx(t *testing.T) {
	t.Run("cancel take", func(t *testing.T) {

		ctxRepeat, cancelRepeat := context.WithCancel(context.Background())
		defer cancelRepeat()
		ch1 := repeatFn(ctxRepeat, fnRand)

		ctxTake, cancelTake := context.WithCancel(context.Background())
		defer cancelTake()
		ch2 := take(ctxTake, ch1, 100)
		for i := 0; i < 50; i++ {
			num, ok := <-ch2
			assert.True(t, ok)
			_, ok = (num).(int)
			assert.True(t, ok)
		}
		cancelTake()
		_, ok := <-ch2 //take the last sent val
		_, ok = <-ch2
		assert.False(t, ok)

		_, ok = <-ch1
		assert.True(t, ok)
		cancelRepeat()
		_, ok = <-ch1 //take the last sent val
		_, ok = <-ch1
		assert.False(t, ok)
	})

	t.Run("cancel repeat", func(t *testing.T) {
		ctxRepeat, cancelRepeat := context.WithCancel(context.Background())
		defer cancelRepeat()
		ch1 := repeatFn(ctxRepeat, fnRand)

		ctxTake, cancelTake := context.WithCancel(context.Background())
		defer cancelTake()
		ch2 := take(ctxTake, ch1, 100)
		for i := 0; i < 50; i++ {
			num, ok := <-ch2
			assert.True(t, ok)
			_, ok = (num).(int)
			assert.True(t, ok)
		}
		cancelRepeat()
		_, ok := <-ch2 //take the last sent val
		_, ok = <-ch2
		assert.False(t, ok)

		_, ok = <-ch1
		assert.False(t, ok)
	})
}
