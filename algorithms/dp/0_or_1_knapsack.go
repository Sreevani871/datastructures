package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		profits = []int{1, 6, 10, 16}
		weights = []int{1, 2, 3, 5}
	)

	fmt.Println("Max profit: ", bruteforceSolution(weights, profits, 7))
	fmt.Println("Max profit: ", bruteforceSolution(weights, profits, 6))

}

func bruteforceSolution(weights []int, profits []int, capacity int) int {
	return knapsackRecursive(weights, profits, capacity, 0)
}

func knapsackRecursive(weights []int, profits []int, capacity int, currentIndex int) int {
	var profit1, profit2 int

	if (capacity == 0) || currentIndex >= len(profits) {
		return 0
	}

	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + knapsackRecursive(weights, profits, capacity-weights[currentIndex], currentIndex+1)
	}

	profit2 = knapsackRecursive(weights, profits, capacity, currentIndex+1)

	maxProfit := math.Max(float64(profit1), float64(profit2))
	return int(maxProfit)
}
