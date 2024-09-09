// Package task_5
//
// You need to write a worker pool:
// it should execute numJobs tasks in parallel, using numWorkers goroutines,
// which are launched once during the program execution.
package task_5

import (
	"fmt"
	"sync"
)

//

// worker:
//
// reads nums from jobs, runs f(job), sends the result to the result chan
func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- f(job)
	}
}

const numJobs = 5
const numWorkers = 3

// main is a function the task asked me to write
//
// runs the worker func in numWorkers goroutines using multiplier as an argument,
// writes nums 1 to numJobs into the jobs chan,
// reads and prints out results from the results chan in parallel
func main() {

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	wg := sync.WaitGroup{}

	multiplier := func(x int) int {
		return x * 10
	}

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			worker(multiplier, jobs, results)
		}()
	}

	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("result: %d\n", result)
	}
}

// runWorkers is how I would really organize the code
//
//	go func() {
//		for i := 1; i <= numJobs; i++ {
//			jobs <- i
//		}
//		close(jobs)
//
// }()
//
// for result := range results {
// fmt.Printf("result: %d\n", result)
// }
func runWorkers(jobs <-chan int, numWorkers int, fn func(x int) int) <-chan int {
	wg := sync.WaitGroup{}
	out := make(chan int)
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			worker(fn, jobs, out)
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func fnMultiplier(x int) int {
	return x * 10
}
