package main

import "fmt"

func Index(s string, toFind string) int {
	var firstIndex int
	if toFind == "" {
		return 0
	}
	if len(toFind) > len(s) {
		return -1
	}
	for i := 0; i < len(s); i++ { // Alternative => for i := range s {}
		if s[i] == toFind[0] { // if any index value of s string equals the first index value of toFind string
			firstIndex = i               // i value assigned to firstIndex; index value equals the firstIndex
			word := s[i : i+len(toFind)] // i has been redefined to firstIndex; accessing values from firstIndex; alternative => word := s[firstIndex : firstIndex+len(toFind)]
			if word == toFind {
				return firstIndex // for i, v := range s {}, v prints rune value but s[i] is byte value
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}

/* Alternative:

package main

import (
"fmt"
"strings"
)

func main() {
fmt.Println(strings.Index("Hello!", "l"))
fmt.Println(strings.Index("Salut!", "alu"))
fmt.Println(strings.Index("Ola!", "hOl"))
}

*/
