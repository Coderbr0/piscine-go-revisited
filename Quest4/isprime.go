package main

import "fmt"

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 { 	// if modulus (remainder) equals zero; remainder of number divided by index equals zero
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(3))
}
