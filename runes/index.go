package runes

import (
	"fmt"
	"unicode/utf8"
)

func Runes() {
	rStr := "abcdefg"
	fmt.Println("Rune count: ", utf8.RuneCountInString(rStr)) //

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal) // Print the index, unicode, character
	}
}
