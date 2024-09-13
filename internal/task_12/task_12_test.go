package task_12

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_groupWait(t *testing.T) {
	g := newGroupWait(2)

	ctx := context.Background()
	expErr1 := errors.New("got error 1")
	expErr2 := errors.New("got error 2")
	g.run(ctx, func(ctx context.Context) error {
		return nil
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr2
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr1
	})

	err := g.wait()
	assert.Equal(t, err, errors.Join(expErr1, expErr2))
}
