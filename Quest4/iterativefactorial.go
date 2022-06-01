package main

import "fmt"

func iterFactorial(nb int) int {
	result := 1

	if nb < 0 {
		return 0
	}

	for nb >= 1 {
		result *= nb
		nb -= 1
	}
	return result
}

func main() {
	arg := 4
	fmt.Println(iterFactorial(arg))
}
/* Alternatives:

package main

import "fmt"

func iterFactorial(nb int) int {
result := 1
if nb < 0 {
return 0
}

if nb <= 1 { // for 0 and 1
return 1
}

for i := 1; i <= nb; i++ {
result = result * i
//fmt.Println(result)
}
return result
}

func main() {
arg := 4
fmt.Println(iterFactorial(arg))
}
----------------------------------------------------
package main

import "fmt"

func iterFactorial(nb int) int {
result := 1
if nb < 0 {
return 0
}

if nb <= 1 { // for 0 and 1
return 1
}

for i := nb; i > 0; i-- {
result = result * i
//fmt.Println(result)
}
return result
}

func main() {
arg := 4
fmt.Println(iterFactorial(arg))
}
----------------------------------------------------
package main

import "fmt"

func iterFactorial(nb int) int {
result := 1
if nb < 0 {
return 0
}

if nb == 0 {
return 1
}

if nb == 1 {
return 1
}

for i := nb + 1; i >= 2; i-- {
result = result * (i - 1)
//fmt.Println(result)
}
return result
}

func main() {
arg := 4
fmt.Println(iterFactorial(arg))
}

//4*3*2*1

1*4 = 4 i=5
4*3 = 12 i=4
12*2 = 24 i=3
24*1 = 24 i=2

*/