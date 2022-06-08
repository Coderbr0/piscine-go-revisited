package main

import "fmt"

func Sqrt(nb int) int {
	for i := 0; i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}

/* Alternative:

package main

import (
"fmt"
"math"
)

func Sqrt(nb int) int {
var x float64 = float64(nb)
result := math.Sqrt(x)
if result*result == float64(nb) {
return int(result)
}
return 0
}

func main() {
fmt.Println(Sqrt(4))
fmt.Println(Sqrt(3))
}

This works since:

func Sqrt(nb int) int {
var x float64 = float64(nb)
result := math.Sqrt(x)
return int(result)
}

func main() {
fmt.Println(Sqrt(4))
fmt.Println(Sqrt(3))
}

Prints:
2
1

result*result == float64(nb) =>
1*1 = 1 (not 3, so 0 is printed out instead)

*/
