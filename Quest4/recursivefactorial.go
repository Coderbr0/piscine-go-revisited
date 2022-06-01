package main

import "fmt"

func recFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb <= 1 {
		return 1
	}
	return nb * recFactorial(nb-1)
}

func main() {
	arg := 4
	fmt.Println(recFactorial(arg))
}

/*

RecursiveFactorial(5)
5 * RecursiveFactorial(4)
5 * 4 * RecursiveFactorial(3)
5 * 4 * 3 * RecursiveFactorial(2)
5 * 4 * 3 * 2 * RecursiveFactorial(1)
5 * 4 * 3 * 2 * 1 = 120

*/
