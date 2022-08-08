package main

import (
	"fmt"
	"unicode"
)

func IsPrintable(s string) bool {
	for _, v := range s {
		if !unicode.IsPrint(v) || !unicode.IsGraphic(v) { // IsPrint categorization is the same as IsGraphic except that the only spacing character is ASCII space, U+0020; IsGraphic covers more spaces (Z).
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}

/*

Letters (L), Marks (M), Numbers (N), Punctuation (P), Symbols (S), and Spaces (Z).

*/

/* Alternative:

func IsPrintable(s string) bool {
	for i, v := range s {
		if v >= 32 && v <= 126 { // Alternative: v >= ' ' && v <= '~'
			i++
		} else {
			return false
		}
	}
	return true
}

*/

// Reference: https://theasciicode.com.ar/ascii-control-characters/delete-ascii-code-127.html
