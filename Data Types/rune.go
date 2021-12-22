package main

import (
	"fmt"
)

func main() {
	s := "GÖ"

	s_rune := []rune(s)
	s_byte := []byte(s)

	fmt.Println(s_rune) // [71 214]
	fmt.Println(s_byte) // [71 195 150]
}
