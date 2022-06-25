package main

import "fmt"

func IsPrimeNum(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(Nb int) int {
	if Nb <= 1 {
		return 2
	}
	var prime int = Nb
	var found bool = false
	// Loop continuously until IsPrimeNum returns true for a number greater than nb
	for !found {
		if IsPrimeNum(prime) { // if IsPrimeNum(prime) == true
			found = true
		} else {
			prime++
			if IsPrimeNum(prime) {
				found = true
			}
		}
	}
	return prime
}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}
