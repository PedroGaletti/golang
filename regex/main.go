package main

import (
	"fmt"
	"regexp"
)

func RegexFuncs() {
	str := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[ˆ ]?", str)
	fmt.Println("match =", match) // match = false

	sStr := "Cat rat mat fat pat"
	r, _ := regexp.Compile("[Crmfp]at")
	fmt.Println("Match string: ", r.MatchString(sStr))                        // Match string:  true
	fmt.Println("Find string: ", r.FindString(sStr))                          // Find string:  Cat
	fmt.Println("Index string: ", r.FindStringIndex(sStr))                    // Index string:  [0 3]
	fmt.Println("All string: ", r.FindAllString(sStr, -1))                    // All string:  [Cat rat mat fat pat]
	fmt.Println("1st 2nd string: ", r.FindAllString(sStr, 2))                 // 1st 2nd string:  [Car rat]
	fmt.Println("All submit index: ", r.FindAllStringSubmatchIndex(sStr, -1)) // All submit index:  [[0 3] [4 7] [8 11] [12 15] [16 19]]
	fmt.Println("Replace all string: ", r.ReplaceAllString(sStr, "Dog"))      // Replace all string:  Dog Dog Dog Dog Dog
}

func main() {
	RegexFuncs()
}
