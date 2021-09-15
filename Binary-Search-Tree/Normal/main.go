// Binary search tree.
package main

import "fmt"

// Binary search tree object structure.
type binarySearchTreeObject struct {
	root *nodeObject
}

// Node object structure defined.
type nodeObject struct {
	value int
	left  *nodeObject
	right *nodeObject
}

// This function insert new value in binary search tree.
func (bst *binarySearchTreeObject) insert(value int) {
	bst.root = put(bst.root, value)
}

// This function put value in node.
func put(node *nodeObject, value int) *nodeObject {
	if node == nil {
		// If node is nil than return new node.
		return &nodeObject{value: value}
	}
	if value > node.value {
		// Right subtree.
		node.right = put(node.right, value)
	} else {
		// Left subtree.
		node.left = put(node.left, value)
	}
	return node
}

// This function search given value in binary search tree.
func (bst binarySearchTreeObject) search(value int) bool {
	var currentNode = bst.root
	for {
		if currentNode == nil {
			// Not found.
			return false
		}
		if currentNode.value == value {
			// Value match found.
			return true
		}
		if value < currentNode.value {
			// Search in left tree.
			currentNode = currentNode.left
		} else {
			// Search in right tree.
			currentNode = currentNode.right
		}
	}
}

// This function show values of binary search tree.
func (bst binarySearchTreeObject) show() {
	inOrder(bst.root)
}

// This function perform in order traversal of binary search tree.
func inOrder(node *nodeObject) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Println(node.value)
	inOrder(node.right)
}

func main() {

	// Binary search tree.
	bst := binarySearchTreeObject{}

	bst.insert(3)
	bst.insert(4)
	bst.insert(9)
	bst.insert(1)
	bst.insert(2)
	bst.show()
	fmt.Println("Search 2 : ", bst.search(2))
	fmt.Println("Search 6 : ", bst.search(6))
}
