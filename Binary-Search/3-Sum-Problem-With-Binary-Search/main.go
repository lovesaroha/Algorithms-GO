// 3 sum problem with binary search.
package main

import "fmt"

// Sorted array.
var numbers = [2000]int{}

// This function performs binary search on given array.
func binarySearch(numbers [2000]int, low int, a int) bool {
	var high = 1999
	for {
		// Divide the array and check if number is greater or lower than the middle.
		var middle = (low + high) / 2
		if numbers[low] == a || numbers[middle] == a || numbers[high] == a {
			// Check if number matches given positions.
			return true
		}
		if low == high || low == middle || high == middle {
			return false
		}
		if a > numbers[middle] {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
}

// This function return 3 unique numbers with sum equal to given value.
func find3SumPairs(numbers [2000]int, sum int) {
	var totalPairs = 0
	for i := 0; i < 2000; i++ {
		for j := i + 1; j < 2000; j++ {
			if numbers[i] >= sum || numbers[j] >= sum || (numbers[i]+numbers[j]) >= sum {
				break
			}
			var third = sum - (numbers[i] + numbers[j])
			if third <= numbers[i] || third <= numbers[j] {
				continue
			}
			if binarySearch(numbers, j, third) {
				// Check if third pair is in array using binary search.
				totalPairs++
			}
		}
	}
	fmt.Println("Total pairs for sum of ", sum, " : ", totalPairs)
}

// This function initialize values.
func assignValues(numbers *[2000]int) {
	for i := 0; i < 2000; i++ {
		numbers[i] = i + 1
	}
}

func main() {
	// Assign values from 1 to 200.
	assignValues(&numbers)

	// Show three sum pairs.
	find3SumPairs(numbers, 1210)
	find3SumPairs(numbers, 51)
	find3SumPairs(numbers, 6)
	find3SumPairs(numbers, 1919)
}
