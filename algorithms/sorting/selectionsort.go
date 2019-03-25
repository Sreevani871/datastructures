package main

/*
Time Complexity: O(n2) as there are two nested loops.

Auxiliary Space: O(1)
The good thing about selection sort is it never makes more than O(n) swaps and can be useful when memory write is a costly operation.

Exercise :
Sort an array of strings using Selection Sort

Stability : The default implementation is not stable. However it can be made stable. Please see stable selection sort for details.

In Place : Yes, it does not require extra space.
*/
import (
	"fmt"
	"time"
)

var (
	temp int
)

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min_index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min_index] > arr[j] {
				min_index = j
			}
		}
		swap(&arr[i], &arr[min_index])
	}
	return arr
}

func swap(i *int, j *int) {
	temp = *i
	*i = *j
	*j = temp
}

func main() {
	var arr []int = []int{10, 3, 2, 7, 8, 1, 10, -1}
	start := time.Now()
	fmt.Println("Unsorted Array: ", arr, start)
	sortedArr := selectionSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))

	arr = []int{1, 2, 3, 4, 5, 6, 7}
	start = time.Now()
	fmt.Println("Unsorted Array: ", arr)
	sortedArr = selectionSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))

}
