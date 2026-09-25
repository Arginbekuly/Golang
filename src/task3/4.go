package main

import "fmt"

func main() {
	inventory := map[string]int{
		"apples":  10,
		"bananas": 5,
	}

	fmt.Println(`Counter of apples before the changing:`, inventory["apples"]) // Output: 10

	inventory["oranges"] = 15

	inventory["apples"] = 12

	delete(inventory, "bananas")

	fmt.Println(inventory)
}
