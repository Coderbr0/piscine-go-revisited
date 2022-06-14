package main

import "fmt"

func FindNextPrime(nb int) int {
	if nb < 2 {
		return false

		for i := 2; i < nb; i++ {
			if nb%i == 0 { 
				return false
			}
		}
	}
}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}