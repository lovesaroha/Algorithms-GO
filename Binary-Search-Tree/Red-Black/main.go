// Red black binary search tree.
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
	red   bool
}

// This function insert new value in binary search tree.
func (bst *binarySearchTreeObject) insert(value int) {
	bst.root = put(bst.root, value)
}

// This function put value in node.
func put(node *nodeObject, value int) *nodeObject {
	if node == nil {
		// If node is nil than return new node.
		return &nodeObject{value: value, red: true}
	}
	if value > node.value {
		// Right subtree.
		node.right = put(node.right, value)
	} else {
		// Left subtree.
		node.left = put(node.left, value)
	}

	// Balance red black binary search tree.
	if isRed(node.right) && !isRed(node.left) {
		// Right node link is red so rotate left.
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		// Two red links in left so rotate right.
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		// Two red links change colors.
		changeColors(node)
	}
	return node
}

// This function return true if given node link is red.
func isRed(node *nodeObject) bool {
	if node == nil {
		return false
	}
	return node.red
}

// This function rotate given node left.
func rotateLeft(node *nodeObject) *nodeObject {
	var rightNode = node.right
	node.right = rightNode.left
	rightNode.left = node
	rightNode.red = node.red
	node.red = true
	return rightNode
}

// This function rotate given node right.
func rotateRight(node *nodeObject) *nodeObject {
	var leftNode = node.left
	node.left = leftNode.right
	leftNode.right = node
	leftNode.red = node.red
	node.red = true
	return leftNode
}

// This function change color of given node.
func changeColors(node *nodeObject) {
	node.left.red = false
	node.right.red = false
	node.red = true
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
