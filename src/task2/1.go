package main

import "fmt"

func main() {

	 a := [5]float64{0.5, 4.5, 6.7, 8.9, 10.1}

	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
