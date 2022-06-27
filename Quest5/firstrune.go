package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	sRune := []rune(s)
	return sRune[0]
}

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

/* Alternative:

func FirstRune(s string) rune {
sRune := []rune(s)
for i := range sRune {
z01.PrintRune(rune(i))
}
return sRune[0]
}

// for loop is unnecessary but code works
*/
