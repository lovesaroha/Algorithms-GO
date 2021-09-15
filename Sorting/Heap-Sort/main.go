// Heap sort algorithm.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This function perform heap sort.
func heapSort(value *[10001]int) {
	for i := (10001 - 1); i > 0; i-- {
		deleteMax(value, i)
	}
}

// This function remove maximum value from the heap.
func deleteMax(values *[10001]int, size int) {
	swap(values, 0, size)
	sink(values, 0, size-1)
}

// This function prepare heap.
func prepareBinaryHeap(values *[10001]int) {
	for i := (10001 - 1) / 2; i >= 0; i-- {
		sink(values, i, (10001 - 1))
	}
}

// This funtion move value down from parent node if it is smaller.
func sink(values *[10001]int, n int, size int) {
	if size == 0 {
		return
	}
	for {
		if (2*n)+1 > size {
			return
		}
		// Find the max child element.
		var max = maximum(*values, (2*n)+1, (2*n)+2, size)
		if values[max] > values[n] {
			// Child node is bigger than parent.
			swap(values, max, n)
			n = max
		} else {
			return
		}
	}
}

// This function perform swap.
func swap(values *[10001]int, i int, j int) {
	var temp = values[i]
	values[i] = values[j]
	values[j] = temp
}

// This function return maximum key of the two value.
func maximum(values [10001]int, i int, j int, size int) int {
	if j >= size {
		return i
	}
	if values[i] > values[j] {
		return i
	}
	return j
}

// This function set random values.
func setRandomValues(values *[10001]int) {
	for i := 0; i < len(values); i++ {
		var random = rand.New(rand.NewSource(int64(i)))
		var value = (random.Int() * 100010000) + 1
		values[i] = value
	}
}

// This function checks if array is sorted.
func isSorted(values [10001]int) bool {
	for i := 0; i < len(values)-1; i++ {
		if values[i+1] < values[i] {
			return false
		}
	}
	return true
}

// This function checks if a array is binary heap.
func isBinaryHeap(values [10001]int) bool {
	for i := 0; i < (10001-1)/2; i++ {
		if (2*i)+1 >= 10001 {
			break
		}
		if values[i] < values[(2*i)+1] || values[i] < values[(2*i)+2] {
			// Check if parent is less than child.
			return false
		}
	}
	return true
}

func main() {
	values := [10001]int{}
	setRandomValues(&values)
	prepareBinaryHeap(&values)
	fmt.Println("Binary heap : ", isBinaryHeap(values))
	// Perform heap sort on random values.
	var start = time.Now()
	heapSort(&values)
	var end = time.Now()
	fmt.Println("Size : 10001 Sorted : ", isSorted(values), " Time : ", end.Sub(start).Nanoseconds(), "ns")
}
