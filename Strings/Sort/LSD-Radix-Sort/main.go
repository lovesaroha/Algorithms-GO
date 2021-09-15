// LSD radix sort.
package main

import "fmt"

// This function sort given string list.
func sort(list [12]string) [12]string {
	var sorted [12]string
	// For every digits from right to left perform key indexed counting.
	for i := 2; i >= 0; i-- {
		// Prepare keys array.
		var keys [27]int
		for j := 0; j < 12; j++ {
			var runes = []rune(list[j])
			keys[int(runes[i])-96] += 1
		}
		// Compute cumulates.
		for j := 0; j < 26; j++ {
			keys[j+1] += keys[j]
		}
		// Sort values in sorted array.
		for j := 0; j < 12; j++ {
			var runes = []rune(list[j])
			var index = int(runes[i]) - 97
			sorted[keys[index]] = list[j]
			keys[index] += 1
		}
	}
	return sorted
}

func main() {
	// String list.
	list := [12]string{"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace"}
	sorted := sort(list)
	fmt.Println(sorted)
}
