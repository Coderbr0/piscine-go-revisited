package main

import (
	"fmt"
	"unicode"
)

func IsLower(s string) bool {
	for _, v := range s {
		if !unicode.IsLower(v) || !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}

/* Alternatives:

func IsLower(s string) bool {
	for i, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			i++
		} else {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	count := 0
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			count++
		} else {
			return false
		}
	}
	return true
}

*/
