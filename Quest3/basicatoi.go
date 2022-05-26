package main

import (
	"fmt"
	"strconv"
)

func BasicAtoi(s string) int {
	y, _ := strconv.Atoi(s) // func Atoi(s string) (int, error) returns 2 values: y is int and _ ignores errors
	return y
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

/* Alternative:

func BasicAtoi(s string) int {
	var result int

	for _, num := range s {
		// fmt.Println(num) => prints rune value
		conversion := int(num) - 48 // casting num (rune value) to int; conversion is now int
		// fmt.Println(result) => prints first value in iteration which is 0 from "var result int"
		// fmt.Println(conversion) => prints int values
		// result = (result * 10) + conversion // 0 + 1 = 1; then 10 + 2 = 12; then 120 + 3 = 123; then 1230 + 4 = 1234; 12340 + 5 = 12345
		result = (result * 10) + conversion
	}
	return result
}

*/
