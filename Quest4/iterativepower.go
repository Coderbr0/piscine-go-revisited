package main

import "fmt"

func IterativePower(nb int, power int) int {
	result := nb * nb * nb
	
	return result
}

func main() {
	fmt.Println(IterativePower(4, 3))
}
