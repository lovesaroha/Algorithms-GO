// Merge sort algorithm.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This function performs merge sort.
func mergeSort(values *[10001]int, temp [10001]int, start int, end int) {
	if end-start == 1 {
		return
	}
	var middle = (start + end) / 2
	// Perform merge sort on two half arrays.
	mergeSort(values, temp, start, middle)
	mergeSort(values, temp, middle, end)
	merge(values, temp, start, end)
}

// This function merge two given array.
func merge(values *[10001]int, temp [10001]int, start int, end int) {
	for i := start; i < end; i++ {
		// Copy values to temp array.
		temp[i] = values[i]
	}
	var middle = (start + end) / 2
	var m = start
	var n = middle
	for i := start; i < end; i++ {
		if m >= middle {
			values[i] = temp[n]
			n++
			continue
		}
		if n >= end {
			values[i] = temp[m]
			m++
			continue
		}
		// Put smaller value in array.
		if temp[m] <= temp[n] {
			values[i] = temp[m]
			m++
		} else {
			values[i] = temp[n]
			n++
		}
	}
}

// This function set random values.
func setRandomValues(values *[10001]int) {
	for i := 0; i < len(values); i++ {
		var random = rand.New(rand.NewSource(int64(i)))
		var value = (random.Int() * 1000) + 1
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

func main() {
	values := [10001]int{}
	setRandomValues(&values)
	var temp [10001]int
	// Perform merge sort on random values.
	var start = time.Now()
	mergeSort(&values, temp, 0, 10001)
	var end = time.Now()
	fmt.Println("Size : 10001 Sorted : ", isSorted(values), " Time : ", end.Sub(start).Nanoseconds(), "ns")
}
