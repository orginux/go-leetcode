package main

import (
	"fmt"
)

func main() {
	// s := "anagram"
	// t := "nagaram"

	s := "aacc"
	t := "ccac"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make(map[rune]int)

	for _, vs := range s {
		letters[vs]++
	}

	for _, vt := range t {
		if _, ok := letters[vt]; !ok {
			return false
		}
	}

	return true
}
