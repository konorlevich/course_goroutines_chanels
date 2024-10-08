package task_1

import "sync"

// fillChan:
// - receives int n
// - returns a chan
// - sends n ints from 0 to n-1 to the chan
func fillChan(n int) (ch chan int) {
	ch = make(chan int)
	go func(ch chan int, n int) {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- i
		}
	}(ch, n)
	return ch
}

// merge
// - receives an array of chans `cs`
// - returns a chan `ch`
// - reads from all the cs channels in parallel and sends ints to the `ch` chan
func merge(cs ...<-chan int) (ch chan int) {
	ch = make(chan int)
	wg := sync.WaitGroup{}
	for _, c := range cs {
		wg.Add(1)
		go func(c <-chan int, ch chan<- int) {
			defer wg.Done()
			for v := range c {
				ch <- v
			}
		}(c, ch)
	}
	go func() {
		defer close(ch)
		wg.Wait()
	}()

	return ch
}
