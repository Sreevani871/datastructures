package main

import (
	"fmt"
)

func binarySearch(arr []int, searchElement int) (found bool) {
	if len(arr) == 0 {
		return found
	}
	middleIndex := len(arr) / 2

	if arr[middleIndex] == searchElement {
		found = true
		return found
	}

	return binarySearch(arr[0:middleIndex], searchElement) || binarySearch(arr[(middleIndex+1):len(arr)], searchElement)
}

func main() {
	var arr = []int{-1, 4, 28, 7, 12, -5, 3, 9, 2, 8, 5}
	for _, e := range arr {
		fmt.Println(binarySearch(arr, e))
	}
	fmt.Println(binarySearch(arr, 120))

}
