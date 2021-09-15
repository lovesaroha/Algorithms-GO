// Stack implementation using linked list.
package main

import "fmt"

// Stack object defined.
type stackObject struct {
	first *nodeObject
}

// Node object structure.
type nodeObject struct {
	value int
	next  *nodeObject
}

// Push function to add value in stack.
func (stack *stackObject) push(value int) {
	var newNode = nodeObject{next: stack.first, value: value}
	stack.first = &newNode
}

// Pop function remove value from stack.
func (stack *stackObject) pop() {
	if stack.first == nil {
		// Empty stack.
		return
	}
	stack.first = stack.first.next
}

// This function shows values of stack.
func (stack stackObject) show() {
	var node = stack.first
	if node == nil {
		// Empty stack.
		return
	}
	for {
		fmt.Println(node.value)
		if node.next == nil {
			return
		}
		node = node.next
	}
}

func main() {

	stack := stackObject{}
	// Add some values in stack.
	stack.push(2)
	stack.push(4)
	stack.push(3)
	stack.push(7)
	stack.pop()
	stack.show()
}
