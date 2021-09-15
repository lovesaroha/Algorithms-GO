// Find substring in a text using brute force.
package main

import "fmt"

// This function return substring starting index.
func findSubstring(text string, substring string) int {
	for i := 0; i <= len(text)-len(substring); i++ {
		// For every character.
		for j := 0; j < len(substring); j++ {
			if substring[j] != text[i+j] {
				break
			}
			if j == len(substring)-1 {
				// Substring found.
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(findSubstring("my name is love", "ame"))
}
