package main // package piscine

import (
	"github.com/01-edu/z01"
)

func main() {	// func PrintComb2() {
	digits := "0123456789"

	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					if a == c && b == d { // Skips 00 00
						continue
					} else if c == a && d < b { // Skips 99 98
						continue
					} else if c < a {	// Skips 99 89
						continue
					} else if (a == c && d > b) || (a != c || b != d) { // 98 99 || 00 11

						z01.PrintRune(rune(digits[a]))
						z01.PrintRune(rune(digits[b]))

						z01.PrintRune(' ')

						z01.PrintRune(rune(digits[c]))
						z01.PrintRune(rune(digits[d]))

						if a == 9 && b == 8 && c == 9 && d == 9 {
							z01.PrintRune('\n')
							break
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
