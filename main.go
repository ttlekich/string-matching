package main

import (
	"fmt"
)

// "aadabcabcdabcdeabcdef"
// "dabab"

func main() {
	str := "ababcababdababababd"
	search := "ababdasdfsdfdf"

	indexes := Find(str, search)

	fmt.Println(indexes)
}

func Find(str string, search string) []int {
	var indexes []int

	lps := NewLPS(search)

	j := 0

	for i := 0; i < len(str); i++ {
		// Consult LPS
		for j > 0 && str[i] != search[j] {
			j = lps[j-1]
		}

		if str[i] == search[j] {
			// Set index if matched search
			if j == len(search)-1 {
				indexes = append(indexes, i-j)
				j = 0
				// Continue on matching search
			} else {
				j++
			}
		}
	}

	return indexes
}

func NewLPS(str string) []int {
	lps := make([]int, len(str))

	j := 0
	lps[0] = 0

	for i := 1; i < len(str); i++ {
		// Consult LPS
		for j > 0 && str[i] != str[j] {
			j = lps[j-1]
		}

		// Set LPS
		if str[i] == str[j] {
			j++
			lps[i] = j
		} else {
			lps[i] = 0
		}

	}

	return lps
}
