package task_6

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func fillChan(ch chan int, ints []int) {
	for _, i := range ints {
		ch <- i
	}
	close(ch)
}

func Test_mergeSorted(t *testing.T) {
	tests := []struct {
		name         string
		intsA, intsB []int
		want         []int
	}{
		{name: "from the task",
			intsA: []int{1, 2, 4}, intsB: []int{-1, 4, 5}, want: []int{-1, 1, 2, 4, 4, 5}},
		{name: "nil",
			intsA: nil, intsB: nil, want: []int{}},
		{name: "empty",
			intsA: []int{}, intsB: []int{}, want: []int{}},
		{name: "all negative, unsorted",
			intsA: []int{-1, -2, -3}, intsB: []int{-5, -6, -7}, want: []int{-5, -6, -7, -1, -2, -3}},
		{name: "all negative, sorted",
			intsA: []int{-3, -2, -1}, intsB: []int{-7, -6, -5}, want: []int{-7, -6, -5, -3, -2, -1}},
		{name: "with doubles",
			intsA: []int{-3, -2, -1}, intsB: []int{-2, -1, 0}, want: []int{-3, -2, -2, -1, -1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := make(chan int)
			b := make(chan int)
			go fillChan(a, tt.intsA)
			go fillChan(b, tt.intsB)

			c := mergeSorted(a, b)
			got := make([]int, 0, len(tt.intsA)+len(tt.intsB))

			for i := range c {
				got = append(got, i)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("wrong sorting: \n%s", diff)
			}
		})
	}
}
