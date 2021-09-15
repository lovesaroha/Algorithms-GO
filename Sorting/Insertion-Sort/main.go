// Inserton sort algorithm.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This function performs insertion sort.
func insertionSort(values *[10000]int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j > 0; j-- {
			// Move smallest to the left.
			if values[j-1] > values[j] {
				var temp = values[j]
				values[j] = values[j-1]
				values[j-1] = temp
			} else {
				break
			}
		}
	}
}

// This function set random values.
func setRandomValues(values *[10000]int) {
	for i := 0; i < len(values); i++ {
		var random = rand.New(rand.NewSource(int64(i)))
		var value = (random.Int() * 1000) + 1
		values[i] = value
	}
}

// This function checks if array is sorted.
func isSorted(values [10000]int) bool {
	for i := 0; i < len(values)-1; i++ {
		if values[i+1] < values[i] {
			return false
		}
	}
	return true
}

func main() {
	values := [10000]int{}
	setRandomValues(&values)
	// Perform insertion sort on random values.
	var start = time.Now()
	insertionSort(&values)
	var end = time.Now()
	fmt.Println("Size : 10000 Sorted : ", isSorted(values), " Time : ", end.Sub(start).Nanoseconds(), "ns")
}
