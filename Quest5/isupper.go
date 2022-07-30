package main

import (
	"fmt"
	"unicode"
)

func IsUpper(s string) bool {
	for _, v := range s {
		fmt.Println(v)
		if !unicode.IsUpper(v) && unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
	fmt.Println(IsUpper("HELlO!"))
}
