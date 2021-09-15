// Stack implementation using array.
package main

import "fmt"

// Stack object defined.
type stackObject struct {
	size   int
	values [10]int
}

// Push function add new value in stack.
func (stack *stackObject) push(value int) {
	stack.size++
	stack.values[stack.size-1] = value
}

// Pop function remove value from stack.
func (stack *stackObject) pop() {
	if stack.size == 0 {
		// Empty stack.
		return
	}
	stack.size--
	stack.values[stack.size] = 0
}

// This function shows stack values.
func (stack stackObject) show() {
	for i := stack.size - 1; i >= 0; i-- {
		fmt.Println(stack.values[i])
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
