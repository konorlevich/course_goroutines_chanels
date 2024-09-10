// Package task_6
//
// Write a function that receives two "sorted" chans and receives a merged sorted chan
package task_6

func mergeSorted(a, b <-chan int) <-chan int {
	out := make(chan int, 2)

	go func() {
		defer close(out)
		var intA, intB int
		var okA, okB bool
		intA, okA = <-a
		intB, okB = <-b
		for {
			if !okA && !okB {
				return
			}
			if okB && (!okA || intB <= intA) {
				out <- intB
				intB, okB = <-b
			}
			if okA && (!okB || intA <= intB) {
				out <- intA
				intA, okA = <-a
			}
		}
	}()

	return out
}
