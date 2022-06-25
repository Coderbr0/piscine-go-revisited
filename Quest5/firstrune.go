package main

import "github.com/01-edu/z01"

func FirstRune(s string) {
	sRune := []rune(s)
	for _, v := range sRune {
		z01.PrintRune(v)
	}
}

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
