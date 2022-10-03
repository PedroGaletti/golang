package main

import (
	"fmt"
	"strings"
)

func Strings() {
	sv1 := "Some word"
	replacer := strings.NewReplacer("Some", "Another")                         // Creating a new replacer func that replace "Some" to "Another"
	sv2 := replacer.Replace(sv1)                                               // Calling Replace func from replacer
	fmt.Println(sv2)                                                           // Print: Another word
	fmt.Println("length: ", len(sv2))                                          // Show length of string: 12
	fmt.Println("Contains another string: ", strings.Contains(sv2, "Another")) // Check if string Contains some value: true
	fmt.Println("o index: ", strings.Index(sv2, "o"))                          // Get the first index of letter in the string: 2
	fmt.Println("Replace: ", strings.Replace(sv2, "o", "0", -1))               // Replace the letter from index: An0ther w0rd

	sv3 := "\nSome Words\n"                                       // \t = tabs \" = double quotes \\ = backslash
	sv3 = strings.TrimSpace(sv3)                                  // Remove the white spaces in the start and in the end
	fmt.Println(sv3)                                              // Print: Some Words
	fmt.Println("Split: ", strings.Split("a-b-c-d", "-"))         // Remove the '-' from the string: [a b c d]
	fmt.Println("Lowercase: ", strings.ToLower(sv3))              // Turn the string into lowercase: some words
	fmt.Println("Lowercase: ", strings.ToUpper(sv3))              // Turn the string into uppercase: SOME WORDS
	fmt.Println("Prefix: ", strings.HasPrefix("tacocat", "taco")) // Check the preffix: true
	fmt.Println("Suffix: ", strings.HasSuffix("tacocat", "cat"))  // Check the suffix: true
}

func main() {
	Strings()
}
