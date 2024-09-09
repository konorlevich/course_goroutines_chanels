package task_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_runWorkers(t *testing.T) {
	tests := []struct {
		name       string
		numWorkers int
		numJobs    int
	}{
		{name: "zero", numWorkers: 0, numJobs: 0},
		{name: "1 worker", numWorkers: 1, numJobs: 0},
		{name: "2 workers, 1 job", numWorkers: 2, numJobs: 1},
		{name: "2 workers, 2 job", numWorkers: 2, numJobs: 2},
		{name: "2 workers, 20000 job", numWorkers: 2, numJobs: 20000},
		{name: "1 workers, 20000 job", numWorkers: 1, numJobs: 20000},
		{name: "1000 workers, 20000 job", numWorkers: 1000, numJobs: 20000},
		{name: "1000 workers, 200000 job", numWorkers: 1000, numJobs: 200000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jobs := make(chan int)

			go func() {
				for i := 1; i <= tt.numJobs; i++ {
					jobs <- i
				}
				close(jobs)

			}()

			out := runWorkers(jobs, tt.numWorkers, fnMultiplier)

			results := make([]int, 0, tt.numJobs)
			for result := range out {
				results = append(results, result)
			}
			assert.Equal(t, tt.numJobs, len(results))
		})
	}
}
