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
