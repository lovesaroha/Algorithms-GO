// Bottom up merger sort algorithm.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This function performs merge sort.
func mergeSort(values *[10001]int, temp [10001]int) {
	for size := 1; size < len(values); size *= 2 {
		// Divide the array in sizes of 2 , 4 , 8 etc.
		for i := 0; i < len(values); i += (2 * size) {
			// Perform merge on small arrays of length  (2 * size).
			var end = minimum(i+(2*size), len(values))
			var middle = minimum(i+size, len(values))
			merge(values, temp, i, middle, end)
		}
	}
}

// This function merge two given array.
func merge(values *[10001]int, temp [10001]int, start int, middle int, end int) {
	for i := start; i < end; i++ {
		// Copy values to temp array.
		temp[i] = values[i]
	}
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

// This function return minimum value.
func minimum(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// This function set random values.
func setRandomValues(values *[10001]int) {
	for i := 0; i < len(values); i++ {
		var random = rand.New(rand.NewSource(int64(i)))
		var value = (random.Int() * 1000100) + 1
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
	mergeSort(&values, temp)
	var end = time.Now()
	fmt.Println("Size : 10001 Sorted : ", isSorted(values), " Time : ", end.Sub(start).Nanoseconds(), "ns")
}
