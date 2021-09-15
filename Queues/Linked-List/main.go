// Queues implementation in linked list.
package main

import "fmt"

// Queue object defined.
type queueObject struct {
	first *nodeObject
	last  *nodeObject
}

// Node object structure.
type nodeObject struct {
	value int
	next  *nodeObject
}

// Enqueue function add new value in queue.
func (queue *queueObject) enqueue(value int) {
	var newNode = nodeObject{value: value}
	if queue.last != nil {
		// Last node next pointer now points to new node.
		queue.last.next = &newNode
	}
	// Last pointer points to new node.
	queue.last = &newNode
	if queue.first == nil {
		queue.first = &newNode
	}
}

// Dequeue function remove value from queue and return the value.
func (queue *queueObject) dequeue() int {
	var value int
	if queue.first == queue.last {
		// Queue have single value.
		queue.last = nil
	}
	value = queue.first.value
	// First points to second node.
	queue.first = queue.first.next
	return value
}

// This function shows values of queue.
func (queue queueObject) show() {
	if queue.first == nil {
		// Empty queue.
		return
	}
	var node = queue.first
	for {
		fmt.Println(node.value)
		if node.next == nil {
			return
		}
		node = node.next
	}
}

func main() {
	queue := queueObject{}

	// Add some values in queue.
	queue.enqueue(3)
	queue.enqueue(7)
	queue.enqueue(11)
	queue.enqueue(9)
	queue.dequeue()
	queue.show()
}
