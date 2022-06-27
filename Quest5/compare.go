package main

import (
	"fmt"
)

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!")) // Capital S comes before small case l in ascii table so "a < b" => -1
	fmt.Println(Compare("Ola!", "Ol"))
}

/* Alternative:

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("Hello!", "Hello!"))
	fmt.Println(strings.Compare("Salut!", "lut!"))
	fmt.Println(strings.Compare("Ola!", "Ol"))
}

*/
