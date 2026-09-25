package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	numbers1 := make([]int, 5)

	var numbers2 = []int{1, 2, 3, 4, 5}

	numbers3 := numbers[0:5]

	fmt.Println("numbers:", numbers)
	fmt.Println("numbers1:", numbers1)
	fmt.Println("numbers2:", numbers2)
	fmt.Println("numbers3:", numbers3)

}
