// Package task_7
package task_7

import (
	"context"
	"sync"
)

type result struct {
	msg string
	err error
}
type search func() *result
type replicas []search

// getFirstResult
// receives context and runs a concurrent search, returns the first available result from replicas
//
// Returns a context error, if the context closed before the first result is available
func getFirstResult(ctx context.Context, replicas replicas) *result {
	if len(replicas) == 0 {
		return nil
	}
	ch := make(chan *result)

	for _, s := range replicas {
		go func(s search) {
			select {
			case <-ctx.Done():
			case ch <- s():
			}
		}(s)
	}

	for {
		select {
		case <-ctx.Done():
			return &result{err: ctx.Err()}
		case r := <-ch:
			return r
		}
	}
}

// getResults
// runs a concurrent search among replicas in replicaKinds, using getFirstResult
//
// returns results for each replicas
func getResults(ctx context.Context, replicaKinds []replicas) []*result {
	if len(replicaKinds) == 0 {
		return nil
	}
	ch := make(chan *result)

	go func() {
		wg := sync.WaitGroup{}
		for _, r := range replicaKinds {
			wg.Add(1)
			go func(r replicas) {
				defer wg.Done()
				ch <- getFirstResult(ctx, r)
			}(r)
		}
		wg.Wait()
		close(ch)
	}()

	results := make([]*result, 0, len(replicaKinds))

	for i := range ch {
		if i != nil {
			results = append(results, i)
		}
	}

	return results
}
