// MSD Radix sort.
package main

import "fmt"

// This function sort given list of string.
func msd(list *[10]string) {
	// Create aux array.
	var aux [10]string
	sort(list, aux, 0, 9, 0)
}

// This function sort strings from left to right.
func sort(list *[10]string, aux [10]string, low int, high int, digit int) {
	if high <= low {
		return
	}
	// One extra key.
	var keys [28]int
	// Prepare keys array.
	for i := low; i < high; i++ {
		var runes = []rune(list[i])
		var index = 0
		if len(runes) > digit {
			index = int(runes[digit]) - 95
		}
		keys[index] += 1
	}
	// Compute cumulates.
	for i := 0; i < 27; i++ {
		keys[i+1] += keys[i]
	}
	// Put values in aux array from low to high.
	for i := low; i <= high; i++ {
		var runes = []rune(list[i])
		var index = 0
		if len(runes) > digit {
			index = int(runes[digit]) - 96
		}
		aux[keys[index]] = list[i]
		keys[index]++
	}
	// Copy values from aux to main array.
	for i := low; i <= high; i++ {
		list[i] = aux[i-low]
	}
	for r := 0; r < 26; r++ {
		// Sort sub array recursively.
		sort(list, aux, low+keys[r], low+keys[r+1]-1, digit+1)
	}
}

func main() {
	// List of string.
	list := [10]string{"string", "sorting", "msd", "algorithm", "implementation", "in", "golang", "by", "love", "saroha"}
	msd(&list)
	fmt.Println(list)
}
