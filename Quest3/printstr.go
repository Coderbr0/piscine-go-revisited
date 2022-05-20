package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	sByte := []byte("Hello World!")
	for _, v := range sByte {
		z01.PrintRune(rune(v))
	}
z01.PrintRune('\n')

}

func main() {
	PrintStr("Hello World!")
}