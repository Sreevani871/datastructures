package main

import (
	"fmt"
)

func Fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}

func main() {
	fmt.Println("Fibnocci series upto 20: ")
	for i := 0; i < 20; i++ {
		fmt.Print(Fib(i), " ")
	}
	fmt.Println()

	for i := 0; i < 20; i++ {
		fmt.Printf("%d factorial value: %d\n", i, Factorial(i))
	}
}
