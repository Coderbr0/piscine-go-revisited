package main

import "fmt"

func Concat(str1 string, str2 string) string {
	s := str1 + str2
	return s
}

func main() {
	fmt.Println(Concat("Hello!", " How are you?"))
}

/* Alternative:

func Concat(str1 string, str2 string) string {
	return str1 + str2
}

*/
