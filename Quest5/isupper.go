package main

import (
	"fmt"
	"unicode"
)

func IsUpper(s string) bool {
	for _, v := range s {
		if !unicode.IsUpper(v) || !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}

/*

if !(unicode.IsUpper(v) || unicode.IsLetter(v)) {} would not work in the case of fmt.Println(IsUpper("hello"))

*/
