package main

import "fmt"

func main() {
	var a = [2][3]int{{1, 2, 3}, {3, 4, 5}}

	fmt.Println(a)

	for _, v := range a {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	for i, v := range a {
		for j, v2 := range v {
			fmt.Printf("a[%d][%d] = %d\n", i, j, v2)
		}
	}


	fmt.Printf("\n")
	for i := range a {
		for j := range a[i] {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
