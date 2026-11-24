package main

import "fmt"

// RepeatAndMissing finds the repeated and the missing value in an array that
// should contain 1..n exactly once but holds one value twice and lacks another.
// XOR of everything collapses to repeated^missing; splitting by its lowest
// set bit separates the two without overflow-prone sums.
func RepeatAndMissing(nums []int) (repeated, missing int) {
	x := 0
	for i, v := range nums {
		x ^= v ^ (i + 1)
	}
	bit := x & -x
	a, b := 0, 0
	for i, v := range nums {
		if v&bit != 0 {
			a ^= v
		} else {
			b ^= v
		}
		if (i+1)&bit != 0 {
			a ^= i + 1
		} else {
			b ^= i + 1
		}
	}
	for _, v := range nums {
		if v == a {
			return a, b
		}
	}
	return b, a
}

func main() {
	r, m := RepeatAndMissing([]int{3, 1, 2, 5, 3})
	fmt.Println("repeated:", r, "missing:", m)
}
