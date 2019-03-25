package main

/*
Time Complexity: O(nlogn) Sorting arrays on different machines. Merge Sort is a recursive algorithm and time complexity can be expressed as following recurrence relation.
T(n) = 2T(n/2) + \Theta(n)
The above recurrence can be solved either using Recurrence Tree method or Master method. It falls in case II of Master Method and solution of the recurrence is \Theta(nLogn).
Time complexity of Merge Sort is \Theta(nLogn) in all 3 cases (worst, average and best) as merge sort always divides the array in two halves and take linear time to merge two halves.

Auxiliary Space: O(n)

Algorithmic Paradigm: Divide and Conquer

Sorting In Place: No in a typical implementation

Stable: Yes

Applications of Merge Sort

	* Merge Sort is useful for sorting linked lists in O(nLogn) time.In case of linked lists the case is different mainly due to difference in memory allocation of arrays and linked lists. Unlike arrays, linked list nodes may not be adjacent in memory. Unlike array, in linked list, we can insert items in the middle in O(1) extra space and O(1) time. Therefore merge operation of merge sort can be implemented without extra space for linked lists.
	In arrays, we can do random access as elements are continuous in memory. Let us say we have an integer (4-byte) array A and let the address of A[0] be x then to access A[i], we can directly access the memory at (x + i*4). Unlike arrays, we can not do random access in linked list. Quick Sort requires a lot of this kind of access. In linked list to access i’th index, we have to travel each and every node from the head to i’th node as we don’t have continuous block of memory. Therefore, the overhead increases for quick sort. Merge sort accesses data sequentially and the need of random access is low.
	* Inversion Count Problem
	* Used in External Sorting


*/
import (
	"fmt"
	"time"
)

var inversionCount = 0

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	m := partition(arr)
	lArr := mergeSort(arr[:m])
	rArr := mergeSort(arr[m:])
	return merge(lArr, rArr)

}

func partition(arr []int) int {
	return len(arr) / 2
}

func merge(lArr []int, rArr []int) []int {
	var (
		n1   = len(lArr)
		n2   = len(rArr)
		size = n1 + n2
		i    = 0
		j    = 0
		arr  = make([]int, size)
	)

	for k := 0; k < size; k++ {

		if i > n1-1 && j <= n2-1 {
			arr[k] = rArr[j]
			j++

		} else if j > n2-1 && i <= n1-1 {
			arr[k] = lArr[i]
			i++
		} else if lArr[i] < rArr[j] {
			arr[k] = lArr[i]
			i++
		} else {
			arr[k] = rArr[j]
			j++
			inversionCount++
		}
	}
	return arr

}

func main() {
	var arr []int = []int{10, 3, 2, 7, 8, 1, 10, -1}
	start := time.Now()
	fmt.Println("Unsorted Array: ", arr, start)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))
	fmt.Println("InversionCount for arr: ", arr, " is ", inversionCount)

	arr = []int{1, 2, 3, 4, 5, 6, 7}
	inversionCount = 0
	start = time.Now()
	fmt.Println("Unsorted Array: ", arr)
	sortedArr = mergeSort(arr)
	fmt.Println("Sorted Array: ", sortedArr, time.Since(start))
	fmt.Println("InversionCount for arr: ", arr, " is ", inversionCount)

}
