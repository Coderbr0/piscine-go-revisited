package main

import "fmt"

func StrRev(s string) string {
	runeOfString := []rune(s)
	var rev []rune // Declare new slice to append to
	for i := len(runeOfString) - 1; i >= 0; i-- {
		rev = append(rev, runeOfString[i])
	}
	// fmt.Println(string(rev))
	return string(rev)
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)	
}

/* Alternative:

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result		// string(v) => Hello World!; first result is "H" and we add this "result" to the next string(v) iteration "e", so we get eH, then leH etc
//  result = result + string(v) would give Hello World!; not the reverse!
}
  return 
}

*/