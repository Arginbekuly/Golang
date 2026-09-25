package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	sl := arr[1:4]

	fmt.Println("Before")
	fmt.Println("arr:", arr)
	fmt.Println("sl:", sl)

	sl[2] = 7

	fmt.Println("\nAfter")
	fmt.Println("arr:", arr)
	fmt.Println("sl:", sl)

}
