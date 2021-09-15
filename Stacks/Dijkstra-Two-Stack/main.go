// Dijkstra two stack implementation using array.
package main

import (
	"fmt"
	"strconv"
)

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
func (stack *stackObject) pop() int {
	if stack.size == 0 {
		// Empty stack.
		return 0
	}
	stack.size--
	return stack.values[stack.size]
}

// This function shows stack values.
func (stack stackObject) show() {
	for i := stack.size - 1; i >= 0; i-- {
		fmt.Println(stack.values[i])
	}
}

func main() {
	valueStack := stackObject{}
	operatorStack := stackObject{}

	var operators = map[string]int{"+": 1, "-": 2, "*": 3, "/": 4}
	expression := "((2+(3*2))*5)"

	for i := 0; i < len(expression); i++ {
		var letter = string(expression[i])
		if letter == "(" {
			continue
		}
		if letter == "+" || letter == "-" || letter == "*" || letter == "/" {
			operatorStack.push(operators[letter])
			continue
		}
		if letter == ")" {
			var operator = operatorStack.pop()
			var value1 = valueStack.pop()
			var value2 = valueStack.pop()
			var result int
			// Perform operation.
			switch operator {
			case 1:
				result = value1 + value2
			case 2:
				result = value2 - value1
			case 3:
				result = value1 * value2
			case 4:
				result = value2 / value1
			}
			valueStack.push(result)
			continue
		}
		value, _ := strconv.Atoi(letter)
		valueStack.push(value)
	}
	valueStack.show()
}
