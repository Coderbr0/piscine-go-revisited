package main

import "fmt"

func AlphaCount(s string) int {
	sRune := []rune(s)
	count := 0
	for i := 0; i < len(sRune); i++ {
		if (sRune[i] >= 65 && sRune[i] <= 90) || (sRune[i] >= 97 && sRune[i] <= 122) { // (rune value >= 65 && <= 90)  (ABC etc) || (rune value >= 97 && <= 122) (abc etc)
			count++
		}
	}
	return count
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

/* Alternative:

func main() {
	s := "Hello 78 World!    4455 /"
	sRune := []rune(s)
	count := 0
	for i := 0; i < len(sRune); i++ {
		if (sRune[i] >= 65 && sRune[i] <= 90) || (sRune[i] >= 97 && sRune[i] <= 122) {
			count++
		}
	}
	fmt.Println(count)
}

*/
