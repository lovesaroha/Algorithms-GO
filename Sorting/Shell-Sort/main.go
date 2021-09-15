// Shell sort algorithm.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This function performs shell sort.
func shellSort(values *[10000]int) {
	// H is a interval or gap.
	for h := len(values) / 2; h > 0; h = h / 2 {
		for i := 0; i < len(values); i++ {
			if (i + h) > len(values)-1 {
				break
			}
			// Similar to insertion sort but with a gap of h.
			for j := i + h; j > 0; j = j - h {
				if j-h < 0 {
					break
				}
				if values[j-h] > values[j] {
					var temp = values[j]
					values[j] = values[j-h]
					values[j-h] = temp
				} else {
					break
				}
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
	shellSort(&values)
	var end = time.Now()
	fmt.Println("Size : 10000 Sorted : ", isSorted(values), " Time : ", end.Sub(start).Nanoseconds(), "ns")
}
