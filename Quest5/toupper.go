package main

import (
	"fmt"
	"strings"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
