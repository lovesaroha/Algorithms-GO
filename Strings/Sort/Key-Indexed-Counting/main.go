// Key indexed counting algorithm to sort characters.
package main

import "fmt"

// This function perform key indexed sorting.
func sort(list [10]string) [10]string {
	var sorted [10]string
	var keys [27]int
	// Prepare keys array.
	for i := 0; i < 10; i++ {
		var runes = []rune(list[i])
		keys[int(runes[0])-96] += 1
	}
	// Compute cumulates.
	for i := 0; i < 26; i++ {
		keys[i+1] += keys[i]
	}
	// Sort values in sorted array.
	for i := 0; i < 10; i++ {
		var runes = []rune(list[i])
		var index = int(runes[0]) - 97
		sorted[keys[index]] = list[i]
		keys[index] += 1
	}
	return sorted
}

func main() {
	// String list.
	list := [10]string{"a", "b", "a", "z", "c", "r", "e", "c", "t", "q"}
	sorted := sort(list)
	fmt.Println(sorted)
}
