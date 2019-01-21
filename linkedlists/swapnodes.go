package main

/*
1) x and y may or may not be adjacent.
2) Either x or y may be a head node.
3) Either x or y may be last node.
4) x and/or y may not be present in linked list.
*/

import (
	"fmt"
)

type List struct {
	Length int
	Start  *Node
	End    *Node
}

var list = &List{}

type Node struct {
	data int
	next *Node
}

func Append(data int) {
	node := &Node{
		data: data,
		next: nil,
	}
	if list.Length == 0 {
		list.Start = node
		list.End = node
	} else {
		list.End.next = node
		list.End = node
	}

	list.Length++
}

func SwapNodesOfData(data1, data2 int) {
	prev1, cur1 := returnNodePrev(data1)
	prev2, cur2 := returnNodePrev(data2)

	if cur1 == nil || cur2 == nil {
		panic("data not found")
		return
	}

	if prev1 == nil {
		list.Start = cur2
	} else {
		prev1.next = cur2
	}

	if prev2 == nil {
		list.Start = cur1
	} else {
		prev2.next = cur1
	}

	temp := cur2.next

	cur2.next = cur1.next

	cur1.next = temp

	if cur1.next == nil {
		list.End = cur1
	}

}

func returnNodePrev(data int) (*Node, *Node) {
	if list.Length == 0 {
		panic("List is empty")
		return nil, nil
	} else if list.Length == 1 {
		panic("List length is 1, swap needs at least list size 2")
		return nil, nil
	} else {
		var previousNode, currentNode *Node
		currentNode = list.Start
		for currentNode.data != data {
			if currentNode.next == nil {
				currentNode = nil
				break
			}
			previousNode = currentNode
			currentNode = currentNode.next
		}

		return previousNode, currentNode
	}
}

func FindMiddleElement() []int {
	var middleElement []int
	middleElementIndex := list.Length / 2
	currentNode := list.Start
	i := 0

	for i < middleElementIndex-1 {
		currentNode = currentNode.next
		i++
	}
	if list.Length%2 == 0 {
		middleElement = append(middleElement, currentNode.data)
		middleElement = append(middleElement, currentNode.next.data)
	} else {
		middleElement = append(middleElement, currentNode.next.data)
	}

	return middleElement

}

func PrintList() {
	if list.Length == 0 {
		panic("list is empty")
		return
	} else {
		var currentNode = list.Start
		for currentNode.next != nil {
			fmt.Print(currentNode.data, " ")
			currentNode = currentNode.next
		}
		fmt.Print(currentNode.data, " ")
	}
	fmt.Println()
}

func main() {
	Append(1)
	Append(2)
	Append(3)
	Append(4)
	Append(5)
	Append(6)

	fmt.Println("Before Swap")
	PrintList()

	fmt.Println("After Swap 2, 3")
	SwapNodesOfData(2, 3)
	PrintList()
	fmt.Println("MiddleElement: ", FindMiddleElement())

	fmt.Println("After Swap 2, 5")
	SwapNodesOfData(2, 5)
	PrintList()
	fmt.Println("MiddleElement: ", FindMiddleElement())

	fmt.Println("After Swap 1, 5")
	SwapNodesOfData(1, 5)
	PrintList()
	fmt.Println("MiddleElement: ", FindMiddleElement())

	fmt.Println("After Swap 5, 6")
	SwapNodesOfData(5, 6)
	PrintList()
	fmt.Println("MiddleElement: ", FindMiddleElement())

	fmt.Println("After Swap 1, 5")
	SwapNodesOfData(1, 5)
	PrintList()
	fmt.Println("MiddleElement: ", FindMiddleElement())

	Append(7)
	fmt.Println("After Swap 1, 5")
	SwapNodesOfData(1, 5)
	PrintList()
	fmt.Println("MiddleElement: ", FindMiddleElement())
}
