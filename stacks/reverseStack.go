package main

import "fmt"

type stack struct {
	data []int // Note: If you define size of array need to handle stack overflow case while adding elements to queue
	top  int
}

var s = stack{}

func Add(item int) {
	s.data = append(s.data, item)
	s.top++
}

func Pop() int {
	if s.top-1 < 0 {
		panic("stack underflow")
	}
	e := s.data[s.top-1]
	s.top--
	s.data = s.data[0:s.top]
	return e
}

func Peep() int {
	if s.top-1 < 0 {
		panic("stack underflow")
	}
	return s.data[s.top-1]
}

func reverse() {
	if s.top > 0 {
		temp := Pop()
		reverse()
		insertAtBottom(temp)
	}

}

func insertAtBottom(item int) {
	if s.top == 0 {
		Add(item)
	} else {
		temp := Pop()
		insertAtBottom(item)
		Add(temp)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		Add(i)
	}

	fmt.Println("Stack: ", s.data)
	reverse()
	fmt.Println("ReverseStack: ", s.data)

	// fmt.Println(Peep())

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(Pop())
	// }

	// fmt.Println(Peep())
	// fmt.Println(Pop())
}
