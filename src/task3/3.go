package main

import "fmt"

func main() {

	scores := map[string]int{
		"Alice": 90,
		"Bob":   0,
	}

	if score, ok := scores["Bob"]; ok {
		fmt.Println("Bob's score:", score)
	} else {
		fmt.Println("Bob's score not found")
	}

	if score, ok := scores["Charlie"]; ok {
		fmt.Println("Charlie's score:", score)
	} else {
		fmt.Println("Charlie's score not found")
	}
}
