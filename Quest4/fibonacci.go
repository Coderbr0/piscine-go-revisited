package main

import "fmt"

func fibonacci(index int) int {
if index < 0 {
return -1
}
if index <= 1 {
return index
}
if index > 1 {
return fibonacci(index - 1) + fibonacci(index - 2)
}
return index
}

func main() {
nthTerm := 4
fmt.Println(fibonacci(nthTerm))
}

/*

0 1 1 2 3 5 8 13 (Fib)

0 1 2 3 4 5 6 7 (Index)

nthTerm = 4 (index 4)

fib(3)+fib(2)
fib(2)+fib(1)+fib(1)+fib(0)
fib(1)+fib(0)+fib(1)+fib(1)+fib(0)
1 + 0 + 1 + 1 + 0 = 3

*/