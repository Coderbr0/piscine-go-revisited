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

/* Alternatives:

func IsUpper(s string) bool {
	for i, letter := range s {
		if letter >= 'A' && letter <= 'Z' {		// Using i++ as a dummy to return true
			i++
		} else {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {
	count := 0
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' {		// Using count++ as a dummy to return true
			count++
		} else {
			return false
		}
	}
	return true
}

*/
