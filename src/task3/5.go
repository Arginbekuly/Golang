package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)

	words := strings.Fields(s)

	for _, word := range words {
		count[word]++
	}
	return count
}

func main() {
	fmt.Println(WordCount("It is ok"))
}
