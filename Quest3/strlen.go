package main

import "fmt"

func StrLen(s string) int {
	length := len([]rune(s))
	return length
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
