package main

import (
	"fmt"
	"strings"
)

var words []string

func main() {
	fmt.Println("Hello world")
}

func SetWords(w []string) {
	words = w
}

func AutoComplete(w string) []string {
	matches := []string{}
	for _, word := range words {
		if strings.HasPrefix(strings.ToLower(word), strings.ToLower(w)) {
			matches = append(matches, word)
		}
	}

	return matches
}
