package main

import (
	"fmt"
	"unicode"
)

func IsAlpha(s string) bool {
	for _, v := range s {
		if unicode.IsMark(v) || unicode.IsPunct(v) || unicode.IsSymbol(v) || unicode.IsSpace(v) { // M, P, S, Z
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}

/*

Letters (L), Marks (M), Numbers (N), Punctuation (P), Symbols (S), and Spaces (Z).

*/

/* Alternative:

func IsAlpha(s string) bool {
for i, v := range s {
if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
i++
} else {
return false
}
}
return true
}

*/
