// Selection sort algorithm.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This function performs selection sort.
func selectionSort(values *[10000]int) {
	for i := 0; i < len(values)-1; i++ {
		// Pick the smallest value in the array and swap it.
		var minimum = i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[minimum] {
				minimum = j
			}
		}
		// Swap values.
		var temp = values[i]
		values[i] = values[minimum]
		values[minimum] = temp
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
	// Perform selection sort on random values.
	var start = time.Now()
	selectionSort(&values)
	var end = time.Now()
	fmt.Println("Size : 10000 Sorted : ", isSorted(values), " Time : ", end.Sub(start).Nanoseconds(), "ns")
}
