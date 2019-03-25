package main

/*
	TimeComplexity: O(nlogn), worst-case: O(n^2)

	Auxiliary Space: O(n)

	Sorting In Place: Yes

	Stable: No (Can be made stable)

	3-Way QuickSort (Dutch National Flag):
		Consider an array which has many redundant elements.
		For example, {1, 4, 2, 4, 2, 4, 1, 2, 4, 1, 2, 2, 2, 2, 4, 1, 4, 4, 4}.
		If 4 is picked as pivot in Simple QuickSort, we fix only one 4 and recursively process remaining occurrences.

		The idea of 3 way QuickSort is to process all occurrences of pivot and is based on Dutch National Flag algorithm.

		In 3 Way QuickSort, an array arr[l..r] is divided in 3 parts:
			a) arr[l..i] elements less than pivot.
			b) arr[i+1..j-1] elements equal to pivot.
			c) arr[j..r] elements greater than pivot.


*/
import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}

func main() {
	var arr []int = []int{10, 3, 2, 7, 8, 1, 10, -1}
	start := time.Now()
	fmt.Println("Unsorted Array: ", arr, start)
	sortedArr := QuickSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))

	arr = []int{1, 2, 3, 4, 5, 6, 7}
	start = time.Now()
	fmt.Println("Unsorted Array: ", arr)
	sortedArr = QuickSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))

}
