package main

/*
TimeComplexity:

Worst case: Θ(n2).
Best case: Θ(n).
Average case for a random array: Θ(n2)
"Almost sorted" case: Θ(n).

In Place : YES
Stable: YES
*/

import (
	"fmt"
	"time"
)

func inserstionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			k := i
			for j := k; j >= 0; j-- {
				if arr[k] < arr[j] {
					arr[k], arr[j] = arr[j], arr[k]
					k = j
				}

			}
		}
	}
	return arr
}

func main() {
	var arr []int = []int{10, 3, 2, 7, 8, 1, 10, -1}
	start := time.Now()
	fmt.Println("Unsorted Array: ", arr, start)
	sortedArr := inserstionSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))

	arr = []int{1, 2, 3, 4, 5, 6, 7}
	start = time.Now()
	fmt.Println("Unsorted Array: ", arr)
	sortedArr = inserstionSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))

	arr = []int{3, 7, 10, 13, 2, 8, 5}
	start = time.Now()
	fmt.Println("Unsorted Array: ", arr)
	sortedArr = inserstionSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))

}
