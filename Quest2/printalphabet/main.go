package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i <= 25; i++ {
		z01.PrintRune(rune(alphabet[i])) // alphabet[i] converts the string to bytes so we cast it to runes for PrintRune function of z01 package to work
	}
	z01.PrintRune('\n')
}

/* Alternatives:

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i := range alphabet {
		z01.PrintRune(rune(alphabet[i]))
	}
	z01.PrintRune('\n')
}

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, v := range alphabet {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

*/
