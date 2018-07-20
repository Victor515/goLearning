package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner := "Yes我爱go"

	// get bytes
	for _, b := range []byte(banner){
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	// decode UTF-8 and convert to Unicode, then put in rune
	for i, ch := range banner{ // ch is a rune
		fmt.Printf("(%d, %X, %c)", i, ch, ch)
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(banner))

	bytes := []byte(banner)
	for len(bytes) > 0{
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size: ]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// get runes
	for i, ch := range []rune(banner){
		fmt.Printf("%d: %c\n", i, ch)
	}
}
