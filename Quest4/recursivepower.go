package main

import "fmt"

func recPower(nb, power int) int {
	if power < 0 {
		return 0
	}
	if power > 0 {
		return nb * recPower(nb, power-1)
	} 
	return 1
}

func main() {
	fmt.Println(recPower(4,3))
}

/*
if power > 0 {
s := nb * recPower(nb, power-1)
fmt.Println(s)
return s 	// So that stored value is printed out in func main()
}
*/