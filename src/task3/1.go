package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//string
	s := "hello"
	fmt.Println(s)

	for i, c := range s {
		fmt.Println(i, string(c))
	}

	fmt.Println(len(s))

	//runes
	runes := []rune(s)
	fmt.Println(len(runes))
	fmt.Printf("%c", runes[1])

	fmt.Printf("\n%d", utf8.RuneCountInString(s))

	//bytes
	str1 := []byte(s)
	str1[0] = 'y'

	fmt.Println(string(str1))

	//trim
	fmt.Println(strings.TrimLeft("stat hello", "jtost"))
	fmt.Println(strings.TrimRight("stat hello", "stot"))
	fmt.Println(strings.TrimSpace(strings.TrimPrefix("stat hello", "stat")))

}
