package main

import (
	"fmt"
	"strconv"
)

func Atoi(s string) int {
	y, _ := strconv.Atoi(s)
	return y
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

/* Alternative:

func Atoi(s string) int {
result := 0
	minusOrPlusResult := 1

	for i, num := range s {
		conv := int(num) - 48
		// fmt.Println(conv)
		if conv >= 0 && conv <= 9 {	// if num >= 48 && num <= 57
			result = result*10 + conv
		} else if conv == -5 && i == 0 { // if num == 43 && i == 0
			minusOrPlusResult = 1 // fmt.Println(Atoi("+1234"))
		} else if conv == -3 && i == 0 { // if num == 45 && i == 0
			minusOrPlusResult = -1 // fmt.Println(Atoi("-1234"))
		} else {
			return 0
		}
	}
	result *= minusOrPlusResult
	return result
}

*/
