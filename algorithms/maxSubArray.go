// Kadane’s Algorithm or max sum of contiguous array

package main

import (
	"fmt"
)

func maxSubArray(a []int) int {

	var (
		max_so_far      int
		max_ending_here int
	)
	for i := 1; i < len(a); i++ {
		max_ending_here = max_ending_here + a[i]
		if max_ending_here < 0 {
			max_ending_here = 0
		} else if max_so_far < max_ending_here {
			max_so_far = max_ending_here
		}
	}

	return max_so_far
}

func main() {
	var a = []int{-2, -3, 4, -1, -2, 1, 5, -3}
	fmt.Println("array: ", a, "maxSubArray sum: ", maxSubArray(a))
}
