package task_3

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generator_readAll(t *testing.T) {
	t.Run("generate", func(t *testing.T) {
		sent := []int{1, 2, 3}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		ch := generator(ctx, sent...)

		t.Run("read", func(t *testing.T) {
			got := make([]int, 0, 3)
			for v := range ch {
				got = append(got, v)
			}

			if diff := cmp.Diff(got, sent); diff != "" {
				t.Errorf("received: %s", diff)
			}
			_, ok := <-ch
			assert.False(t, ok)
		})
	})
}

func Test_generator_readPart(t *testing.T) {
	t.Run("generate", func(t *testing.T) {
		sent := []int{1, 2, 3, 5}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		ch := generator(ctx, sent...)

		t.Run("read", func(t *testing.T) {
			got := make([]int, 0, 2)
			got = append(got, <-ch, <-ch)
			cancel()

			assert.Equal(t, got, []int{1, 2})
			lastN, ok := <-ch
			lastN, ok = <-ch
			assert.False(t, ok)
			assert.Equal(t, 0, lastN)
		})
	})
}

func Test_squarer_readAll(t *testing.T) {
	t.Run("generate", func(t *testing.T) {
		sent := []int{1, 2, 3}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		generated := generator(ctx, sent...)

		t.Run("read", func(t *testing.T) {
			squared := squarer(ctx, generated)
			got := make([]int, 0, 3)
			for v := range squared {
				got = append(got, v)
			}

			if diff := cmp.Diff(got, []int{1, 4, 9}); diff != "" {
				t.Errorf("received: %s", diff)
			}
			_, ok := <-generated
			assert.False(t, ok)
			_, ok = <-squared
			assert.False(t, ok)
		})
	})
}

func Test_squarer_readPart(t *testing.T) {
	t.Run("generate", func(t *testing.T) {
		sent := []int{1, 2, 3, 5}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		generated := generator(ctx, sent...)

		t.Run("read", func(t *testing.T) {
			got := make([]int, 0, 2)
			squared := squarer(ctx, generated)
			got = append(got, <-squared, <-squared)
			cancel()

			if diff := cmp.Diff(got, []int{1, 4}); diff != "" {
				t.Errorf("received: %s", diff)

			}
			// the last num left in chanel
			lastN, ok := <-generated
			lastN, ok = <-squared
			assert.False(t, ok)
			assert.Equal(t, 0, lastN)
		})
	})
}
