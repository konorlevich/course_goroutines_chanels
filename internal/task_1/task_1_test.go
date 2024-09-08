package task_1

import (
	"github.com/google/go-cmp/cmp"
	"slices"
	"testing"
)

func Test_Task1(t *testing.T) {
	cases := []struct {
		name     string
		sendInts []int
		expect   []int
	}{
		{"one sender - 1", []int{}, []int{}},
		{"one sender - 1", []int{1}, []int{0}},
		{"senders: 2,3,4", []int{1, 2, 3, 4}, []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 3}},
		{"senders: 2,3,4", []int{4, 3, 2, 1}, []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 3}},
	}
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			chans := make([]<-chan int, 0, len(s.sendInts))
			for _, n := range s.sendInts {
				chans = append(chans, fillChan(n))
			}

			d := merge(chans...)

			multiplier := 0
			if len(s.sendInts) != 0 {
				multiplier = slices.Max(s.sendInts)
			}
			res := make([]int, 0, len(s.sendInts)*multiplier)
			for v := range d {
				res = append(res, v)
			}

			t.Run("check ints", func(t *testing.T) {
				if len(res) != len(s.expect) {
					t.Errorf("want 9, got %d", len(res))
				}
				slices.Sort(res)

				if diff := cmp.Diff(s.expect, res); diff != "" {
					t.Errorf("wrong result (-want +got):\n%s", diff)
				}
			})
		})
	}
}
