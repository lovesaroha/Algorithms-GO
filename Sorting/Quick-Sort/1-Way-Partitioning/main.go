// Quick sort algorithm.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This function performs quick sort.
func quickSort(values *[10001]int, start int, end int) {
	if start >= end {
		return
	}
	// Choose last element as a pivot and create partition.
	var p = partition(values, start, end)
	quickSort(values, start, p-1)
	quickSort(values, p, end)
}

// This function perform partition.
func partition(values *[10001]int, start int, end int) int {
	var p = end
	var j = end - 1
	var i = start
	for {
		if i >= j {
			// Partition is complete swap the pivot element.
			if values[i] >= values[p] {
				swap(values, i, p)
				return i
			}
		}
		if values[i] < values[p] {
			// Value in left side already smaller than pivot.
			i++
		} else if values[j] > values[p] {
			// Value in right side greater than pivot.
			j--
		} else {
			if values[j] < values[i] {
				swap(values, i, j)
				i++
				j--
			}
		}
	}
}

// This function perform swap.
func swap(values *[10001]int, i int, j int) {
	var temp = values[i]
	values[i] = values[j]
	values[j] = temp
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
	// Perform quick sort on random values.
	var start = time.Now()
	quickSort(&values, 0, 10000)
	var end = time.Now()
	fmt.Println("Size : 10001 Sorted : ", isSorted(values), " Time : ", end.Sub(start).Nanoseconds(), "ns")
}
