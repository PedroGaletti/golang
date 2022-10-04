package main

import (
	"fmt"
	"regexp"
)

func IsEmail(s string) (string, error) {
	r, _ := regexp.Compile(`[\w._%+-]{1,20}@[\w.-]{2,20}.[A-Za-z]{2,3}`)

	if r.MatchString(s) {
		return "Valid", nil
	}

	return "", fmt.Errorf("not valid")
}

func main() {
	IsEmail("pedro@gmail.com")
}
