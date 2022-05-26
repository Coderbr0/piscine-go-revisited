package main

import (
	"fmt"
	"strconv"
)

func BasicAtoi2(s string) int {
	y, _ := strconv.Atoi(s)
	return y
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

/* Alternative:

func BasicAtoi2(s string) int {
	var result int
	for _, num := range s {
		if num >= 48 && num <= 57 {
		conversion := num - 48
		result = (result * 10) + int(conversion)
		} else {
			return 0
		}
	}
	return result
}

*/
