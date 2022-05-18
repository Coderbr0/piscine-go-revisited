package main

import "fmt"

func PointOne(n *int) {
	*n = 1
}

func main() {
	var p int
	PointOne(&p)
	fmt.Println(p)
}
