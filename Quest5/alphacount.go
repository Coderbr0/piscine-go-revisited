package main

import "fmt"

func main() {
	s := "Hello World"
	count := 0
	for i := 0; i < len(s); i++ {
		fmt.Println(count)
	}
}

/*
rune value >= 65 && <= 90  (ABC etc)
rune value >= 97 && <= 122 (abc etc)
*/
