package task_7

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func fakeSearch(kind string) search {
	return func() *result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return &result{
			msg: fmt.Sprintf("%q result", kind),
		}
	}
}

func Test_getFirstResult(t *testing.T) {
	cases := []struct {
		name               string
		timeout            time.Duration
		replicas           replicas
		wantError          error
		expectedMsgContain string
	}{
		{name: "short timeout",
			timeout: time.Nanosecond, replicas: replicas{fakeSearch("web1"), fakeSearch("web2")},
			wantError: context.DeadlineExceeded},
		{name: "long timeout",
			timeout: 5 * time.Minute, replicas: replicas{fakeSearch("web1"), fakeSearch("web2")},
			expectedMsgContain: "result",
		},
		{name: "no replicas",
			timeout: 5 * time.Minute, replicas: replicas{},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := context.WithTimeout(context.Background(), tt.timeout)
			got := getFirstResult(ctx, tt.replicas)

			if tt.wantError == nil && tt.expectedMsgContain == "" {
				assert.Nil(t, got)
				return
			}

			if tt.wantError != nil {
				assert.ErrorIs(t, got.err, tt.wantError)
			} else {
				assert.NoError(t, got.err)
			}

			assert.Contains(t, got.msg, tt.expectedMsgContain)
		})
	}
}

func Test_getResults_ShortContext(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Nanosecond)
	replicaKinds := []replicas{
		{fakeSearch("web1"), fakeSearch("web2")},
		{fakeSearch("image1"), fakeSearch("image2")},
		{fakeSearch("video1"), fakeSearch("video2")},
	}
	results := getResults(ctx, replicaKinds)
	assert.Len(t, results, 3)
	for _, res := range results {
		assert.ErrorIs(t, context.DeadlineExceeded, res.err)
	}
}

func Test_getResults_EmptyReplicas(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Nanosecond)
	replicaKinds := []replicas{{}, {}, {}}
	got := getResults(ctx, replicaKinds)
	assert.NotNil(t, got)
	assert.Len(t, got, 0)
}

func Test_getResults_EmptyReplicaSet(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	replicaKinds := []replicas{}
	assert.Nil(t, getResults(ctx, replicaKinds))
}

func Test_getResults_LongContext(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	replicaKinds := []replicas{
		{fakeSearch("web1"), fakeSearch("web2")},
		{fakeSearch("image1"), fakeSearch("image2")},
		{fakeSearch("video1"), fakeSearch("video2")},
	}
	results := getResults(ctx, replicaKinds)
	assert.Len(t, results, 3)
	for _, res := range results {
		assert.Nil(t, res.err)
		assert.Contains(t, res.msg, "result")
	}
}
