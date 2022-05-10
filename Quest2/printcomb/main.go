package main // package piscine

import "github.com/01-edu/z01"

func main() { // func PrintComb() {}
	for a := 48; a < 56; a++ {			// a starts at 48 as rune character value for this integer is 0; refer to ascii table (man ascii in terminalc)
		for b := a + 1; b < 57; b++ {
			for c := b + 1; c < 58; c++ {
				if a == 55 && b == 56 && c == 57 {
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(c))
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(c))
					z01.PrintRune(44) // Comma in ascii table
					z01.PrintRune(32) // Space in ascii table
				}
			}
		}
	}
}