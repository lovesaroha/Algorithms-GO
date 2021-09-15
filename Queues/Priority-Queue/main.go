// Unordered priority queue.
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

// This dequeue function remove maximum value from queue.
func (queue *queueObject) dequeueMax() {
	if queue.first == queue.last {
		// Empty queue.
		return
	}
	var max = 0
	for i := queue.first; i < queue.last; i++ {
		// Find the maximum value in queue.
		if queue.values[i] > queue.values[max] {
			max = i
		}
	}
	swap(&queue.values, queue.first, max)
	queue.values[queue.first] = 0
	queue.first++
}

// This dequeue function remove minimum value from queue.
func (queue *queueObject) dequeueMin() {
	if queue.first == queue.last {
		// Empty queue.
		return
	}
	var min = queue.first
	for i := queue.first; i < queue.last; i++ {
		// Find the maximum value in queue.
		if queue.values[i] < queue.values[min] {
			min = i
		}
	}
	swap(&queue.values, queue.first, min)
	queue.values[queue.first] = 0
	queue.first++
}

// This function shows values.
func (queue queueObject) show() {
	for i := queue.first; i < queue.last; i++ {
		fmt.Println(queue.values[i])
	}
}

// This function perform swap.
func swap(values *[10]int, i int, j int) {
	var temp = values[i]
	values[i] = values[j]
	values[j] = temp
}

func main() {
	queue := queueObject{}

	// Add some values in queue.
	queue.enqueue(3)
	queue.enqueue(7)
	queue.enqueue(11)
	queue.enqueue(2)
	queue.enqueue(9)
	queue.dequeueMax()
	queue.dequeueMin()
	queue.show()
}
