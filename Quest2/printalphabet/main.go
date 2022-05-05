package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i <= 25; i++ {
		z01.PrintRune(rune(alphabet[i]))
	}
	z01.PrintRune('\n')
}

// alphabet[i] converts the string to bytes so we cast it to runes for PrintRune function of z01 package to work
