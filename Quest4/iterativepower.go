package main

import (
	"fmt"
	"math"
)

func iterativePower(nb, power int) int {
	if power < 0 {
		return 0
	}
	var x float64 = float64(nb)
	var y float64 = float64(power)
	result := math.Pow(x, y)
	return int(result)
}

func main() {
	fmt.Println(iterativePower(4, 3))
}
