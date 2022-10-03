package main

import (
	"fmt"
	"unicode/utf8"
)

// rune in Go is a data type that stores codes that represent Unicode characters

func Runes() {
	rStr := "abcdefg"
	fmt.Println("Rune count: ", utf8.RuneCountInString(rStr)) // 7

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal) // Print the index, unicode, character
	}
}

func main() {
	Runes()
}
