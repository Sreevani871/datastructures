/*
Worst and Average Case Time Complexity: O(n*n). Worst case occurs when array is reverse sorted.

Best Case Time Complexity: O(n). Best case occurs when array is already sorted.

Auxiliary Space: O(1)

Boundary Cases: Bubble sort takes minimum time (Order of n) when elements are already sorted.

Sorting In Place: Yes

Stable: Yes
*/
package main

import (
	"fmt"
	"strings"
	"time"
)

var (
	temp int
)

func bubleSort(arr []int) []int {
	isSwapped := true
	for i := 0; i < len(arr)-1 && isSwapped; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			isSwapped = false
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
				isSwapped = true
			}
		}
	}
	return arr
	strings.Compare(a, b)
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
	sortedArr := bubleSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))

	arr = []int{1, 2, 3, 4, 5, 6, 7}
	start = time.Now()
	fmt.Println("Unsorted Array: ", arr)
	sortedArr = bubleSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))

}
