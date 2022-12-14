package main

import (
	"fmt"
	"unicode/utf8"
)

// rune in Go is a data type that stores codes that represent Unicode characters

func Runes() {
	rStr := "abcdefg"
	fmt.Println("Rune count: ", utf8.RuneCountInString(rStr)) // Rune count: 7

	for i, runeVal := range rStr {
		fmt.Printf("%d - %#U - %c\n", i, runeVal, runeVal) // Print the index, unicode, character: 0 - U+0061 'a' - a
	}
}

func main() {
	Runes()
}
