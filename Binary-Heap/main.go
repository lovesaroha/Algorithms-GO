// Binary heap implementation using array.
package main

import "fmt"

// Binary heap object structure.
type binaryHeapObject struct {
	values [10]int
	size   int
}

// This function insert value in binary heap.
func (bh *binaryHeapObject) insert(value int) {
	bh.size += 1
	bh.values[bh.size] = value
	bh.swim(bh.size)
}

// This function remove maximum value from the heap.
func (bh *binaryHeapObject) deleteMax() {
	bh.swap(1, bh.size)
	bh.values[bh.size] = 0
	bh.size += -1
	bh.sink()
}

// This function move values up if parent is smaller.
func (bh *binaryHeapObject) swim(n int) {
	for {
		if n == 1 {
			return
		}
		parent := bh.values[n/2]
		if parent < bh.values[n] {
			// Swap the values if parent is smaller.
			bh.swap(n/2, n)
			n = n / 2
		} else {
			return
		}
	}
}

// This function move value down from first node if it is smaller.
func (bh *binaryHeapObject) sink() {
	if bh.size <= 1 {
		// Not enough values in array.
		return
	}
	var n = 1
	for {
		if 2*n > bh.size {
			// No child nodes.
			return
		}
		// Find the max child element.
		var max = bh.maximum(2*n, 2*n+1)
		if bh.values[max] > bh.values[n] {
			// Child node is bigger than parent.
			bh.swap(max, n)
			n = max
		} else {
			return
		}
	}
}

// This function perform swap.
func (bh *binaryHeapObject) swap(i int, j int) {
	var temp = bh.values[i]
	bh.values[i] = bh.values[j]
	bh.values[j] = temp
}

// This function return maximum key of the two value.
func (bh binaryHeapObject) maximum(i int, j int) int {
	if bh.values[i] > bh.values[j] {
		return i
	}
	return j
}

// This function shows values of binary heap.
func (bh binaryHeapObject) show() {
	for i := 1; i < bh.size+1; i++ {
		fmt.Println(bh.values[i])
	}
}

func main() {
	bh := binaryHeapObject{}

	// Add some values in binary heap.
	bh.insert(2)
	bh.insert(4)
	bh.insert(9)
	bh.insert(12)
	bh.insert(3)
	bh.insert(6)
	bh.deleteMax()
	bh.deleteMax()
	bh.show()
}
