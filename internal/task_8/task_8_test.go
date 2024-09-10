package task_8

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_executeTaskWithTimeout(t *testing.T) {
	tests := []struct {
		name       string
		timeout    time.Duration
		ctxTimeout time.Duration
		wantErr    error
	}{
		{
			name:       "ctx shorter",
			timeout:    2 * time.Second,
			ctxTimeout: time.Nanosecond,
			wantErr:    context.DeadlineExceeded,
		},
		{
			name:       "timeout shorter",
			timeout:    time.Nanosecond,
			ctxTimeout: 2 * time.Second,
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeout = tt.timeout
			ctx, cancel := context.WithTimeout(context.Background(), tt.ctxTimeout)
			defer cancel()
			if tt.wantErr != nil {
				assert.ErrorIs(t, executeTaskWithTimeout(ctx), tt.wantErr)
				return
			}
			assert.Nil(t, executeTaskWithTimeout(ctx))
		})
	}
}
