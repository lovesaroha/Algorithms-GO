// Binary search algorithm.
package main

import "fmt"

// Sorted array.
var numbers = [200]int{}

// This function performs binary search on given array.
func binarySearch(numbers [200]int, a int) bool {
	var low = 0
	var high = 199
	for {
		// Divide the array and check if number is greater or lower than the middle.
		var middle = (low + high) / 2
		if numbers[low] == a || numbers[middle] == a || numbers[high] == a {
			// Check if number matches given positions.
			return true
		}
		if a > numbers[middle] {
			low = middle + 1
		} else {
			high = middle - 1
		}
		if low == high {
			return false
		}
	}
}

// This function initialize values.
func assignValues(numbers *[200]int) {
	for i := 0; i < 200; i++ {
		numbers[i] = i + 1
	}
}

func main() {
	// Assign values from 1 to 200.
	assignValues(&numbers)

	// Perform binary search.
	fmt.Println("For 11 : ", binarySearch(numbers, 11))
	fmt.Println("For 86 : ", binarySearch(numbers, 86))
	fmt.Println("For 191 : ", binarySearch(numbers, 191))
	fmt.Println("For 400 : ", binarySearch(numbers, 400))
}
