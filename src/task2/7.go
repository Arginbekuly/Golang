package main

import "fmt"

func main() {
	arr := []int{8, 44, 3, 4, 5}

	dst := make([]int, 2)

	num := copy(dst, arr)

	fmt.Println("My slice:", arr)

	dst[0] = 7

	fmt.Println("Destination slice:", dst)
	fmt.Println("Number of elements copied:", num)
}
