package main

import (
	"fmt"
)

/*
      tree
      ----
       1    <-- root
     /   \
    2     4
   /
  3

*/
type Tree struct {
	root *Node
}

type Node struct {
	data       int
	leftChild  *Node
	rightChild *Node
}

func getNode(data int) *Node {
	return &Node{
		data:       data,
		leftChild:  nil,
		rightChild: nil,
	}
}

func inOrderTraversal(node *Node, f func(int)) {
	if node == nil {
		return
	}

	inOrderTraversal(node.leftChild, f)
	f(node.data)
	inOrderTraversal(node.rightChild, f)

}

func preOrderTraversal(node *Node, f func(int)) {
	if node == nil {
		return
	}
	f(node.data)
	preOrderTraversal(node.leftChild, f)
	preOrderTraversal(node.rightChild, f)
}

func postOrderTraversal(node *Node, f func(int)) {
	if node == nil {
		return
	}
	postOrderTraversal(node.leftChild, f)
	postOrderTraversal(node.rightChild, f)
	f(node.data)
}

func levelOrderTraversal(node *Node) {
	for level := 1; level <= heightOfTree(node); level++ {
		printGivenLevel(node, level)
		fmt.Println()
	}
}

func printGivenLevel(node *Node, level int) {
	if node == nil {
		return
	}
	if level == 1 {
		fmt.Print(node.data, " ")
	} else {
		printGivenLevel(node.leftChild, level-1)
		printGivenLevel(node.rightChild, level-1)
	}
}

func heightOfTree(node *Node) int {
	if node == nil {
		return 0
	} else {
		lh := heightOfTree(node.leftChild)
		rh := heightOfTree(node.rightChild)
		if lh > rh {
			return lh + 1
		} else {
			return rh + 1
		}
	}
}

func isBST(root *Node) bool {

	return isBSTUtil(root, nil, nil)
}

func isBSTUtil(root *Node, l *Node, r *Node) bool {
	if root == nil {
		return true
	}

	if l != nil && root.data < l.data {
		return false
	}

	if r != nil && root.data > r.data {
		return false
	}

	return isBSTUtil(root.leftChild, l, root) && isBSTUtil(root.rightChild, root, r)
}

func printLeftViewOfBinaryTree(root *Node) {
	if root == nil {
		return
	}
	if root.leftChild != nil {
		fmt.Print(root.leftChild.data, " ")
	}
	printLeftViewOfBinaryTree(root.leftChild)
	printLeftViewOfBinaryTree(root.rightChild)

}

func main() {
	var t = &Tree{}
	data := 4
	t.root = getNode(data)
	data = 2
	t.root.leftChild = getNode(data)
	data = 5
	t.root.rightChild = getNode(data)
	data = 6
	t.root.leftChild.leftChild = getNode(data)
	data = 2
	t.root.leftChild.rightChild = getNode(data)
	data = 4
	t.root.leftChild.rightChild.leftChild = getNode(data)

	// fmt.Println(t.root.data, t.root.leftChild.data, t.root.rightChild.data, t.root.leftChild.leftChild.data)

	// Binary tree traversal
	// * DFS - 1.Inorder -> left-root-right
	// 		 2.Preorder-> root-left-right
	//		 3.Postorder-> right-left-root
	// 1. Inorder
	var result = []int{}
	inOrderTraversal(t.root, func(data int) {
		result = append(result, data)
	})

	fmt.Println("InorderTraversal: ", result)

	// 2. Preorder
	result = []int{}
	preOrderTraversal(t.root, func(data int) {
		result = append(result, data)
	})

	fmt.Println("PreOrderTraversal: ", result)

	// 2. Preorder
	result = []int{}
	postOrderTraversal(t.root, func(data int) {
		result = append(result, data)
	})

	fmt.Println("PostOrderTraversal: ", result)

	// * BFS - Level Order
	fmt.Println("LevelOrderTraversal: ", result)
	levelOrderTraversal(t.root)

	fmt.Println("IsBST: ", isBST(t.root))

	fmt.Println("LeftViewOfBT: ")
	printLeftViewOfBinaryTree(t.root)
	fmt.Println()
}
