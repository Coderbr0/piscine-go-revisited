package main

import (
	"fmt"
	"unicode"
)

func IsNumeric(s string) bool {
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsMark(v) || unicode.IsPunct(v) || unicode.IsSymbol(v) || unicode.IsSpace(v) { // L, M, P, S, Z
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}

/*

Letters (L), Marks (M), Numbers (N), Punctuation (P), Symbols (S), and Spaces (Z).

*/

/* Alternative:

func IsNumeric(s string) bool {
	for i, v := range s {
		if v >= '0' && v <= '9' {
			i++
		} else {
			return false
		}
	}
	return true
}

*/
