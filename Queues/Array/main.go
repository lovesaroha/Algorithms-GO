// Queue implementation usign array.
package main

import "fmt"

// Queue structure defined.
type queueObject struct {
	first  int
	last   int
	values [10]int
}

// Enqueue function add new value in queue.
func (queue *queueObject) enqueue(value int) {
	queue.values[queue.last] = value
	queue.last++
}

// Dequeue function remove value from queue.
func (queue *queueObject) dequeue() {
	if queue.first == queue.last {
		// Empty queue.
		return
	}
	queue.values[queue.first] = 0
	queue.first++
}

// This function shows values.
func (queue queueObject) show() {
	for i := queue.first; i < queue.last; i++ {
		fmt.Println(queue.values[i])
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
	queue.dequeue()
	queue.show()
}
