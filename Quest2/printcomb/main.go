package main

import "github.com/01-edu/z01"

func main() {
	for a := '7'; b := '8'; c := '9' {
		if a < b && b < c {
			z01.PrintRune(rune(a))
			z01.PrintRune(rune(b))
			z01.PrintRune(rune(c))
			z01.PrintRune('\n')
		}
	}
}