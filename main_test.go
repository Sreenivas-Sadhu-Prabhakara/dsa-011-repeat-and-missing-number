package main

import "testing"

func TestRepeatAndMissing(t *testing.T) {
	cases := []struct {
		in        []int
		rep, miss int
	}{
		{[]int{3, 1, 2, 5, 3}, 3, 4},
		{[]int{1, 1}, 1, 2},
		{[]int{2, 2}, 2, 1},
		{[]int{4, 3, 6, 2, 1, 1}, 1, 5},
		{[]int{1, 2, 3, 4, 4}, 4, 5},
	}
	for _, c := range cases {
		r, m := RepeatAndMissing(c.in)
		if r != c.rep || m != c.miss {
			t.Errorf("RepeatAndMissing(%v) = (%d,%d), want (%d,%d)", c.in, r, m, c.rep, c.miss)
		}
	}
}
