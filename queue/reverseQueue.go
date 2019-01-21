package main

import (
	"fmt"
)

type queue struct {
	arr   []int
	front int
	rare  int
	size  int
}

var size = 20
var q = queue{
	arr:   make([]int, size),
	front: 0,
	rare:  0,
	size:  size,
}

func enqueue(data int) {
	if q.rare == q.size && q.rare > 0 {
		panic("queue is full")
	}

	q.arr[q.rare] = data
	q.rare = q.rare + 1
}

func dequeue() int {
	if q.front < 0 || (q.front == q.size && q.size > 0) {
		panic("queue is empty")
	}

	de := q.arr[q.front]

	q.front = (q.front + 1) % q.size
	q.rare = q.rare % q.size

	return de
}

func main() {
	for i := 0; i < 20; i++ {
		enqueue(i)
	}
	fmt.Println("Queue: ", q.arr[q.front:q.rare])
	// enqueue(10)

	// for i := 0; i < 10; i++ {
	// 	dequeue()
	// 	if i == 5 {
	// 		fmt.Println("Queue: ", q.arr[q.front:q.rare])
	// 	}
	// }

	// fmt.Println("Queue: ", q.arr[q.front:q.rare])
	for i := 0; i < 20; i++ {
		e := dequeue()
		push(e)
	}
	for i := 0; i < 20; i++ {
		e := pop()
		enqueue(e)
	}
	fmt.Println("ReverseQueue: ", q.arr[q.front:q.rare])

}

type stack struct {
	arr  []int
	top  int
	size int
}

var s = stack{
	arr:  make([]int, size),
	top:  0,
	size: size,
}

func push(data int) {
	if s.top == s.size && s.size > 0 {
		panic("stack overflow")
	}

	s.arr[s.top] = data
	if s.top+1 < s.size {
		s.top++
	}
}

func pop() int {
	if s.top < 0 {
		panic("stack underflow")
	}

	e := s.arr[s.top]
	s.top--
	return e
}
