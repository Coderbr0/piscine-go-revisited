package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "zyxwvutsrqponmlkjihgfedcba"

	for _, v := range alphabet {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

/* Alternatives:

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i := 25; i >= 0; i-- {
		z01.PrintRune(rune(alphabet[i])) // Casting is not allowed for this exercise!
	}
	z01.PrintRune('\n')
}

*/
