package main

import (
"fmt"
)

func IsUpper(s string) bool {
runeTheString := []rune(s)
fmt.Println(runeTheString)
for i := 0; i < len(runeTheString); i++ {
if (runeTheString[i] >= 65 && runeTheString[i] <= 90) { // (rune value >= 65 && <= 90)  (ABC etc)
return true
}
}
return false
}

func main() {
fmt.Println(IsUpper("HELLO"))
fmt.Println(IsUpper("HELLO!"))
fmt.Println(IsUpper("HeLLO!"))
}

/* Note: Prints true at any value found within the range; if any value is greater than 65 and less than 90 */