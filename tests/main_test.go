package main

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("Hello is not an email")
	}

	_, err = IsEmail("pedro@gmail.com")
	if err == nil {
		t.Error("pedro@gmail.com is an email")
	}

	_, err = IsEmail("pedro@gmail")
	if err == nil {
		t.Error("pedro@gmail is not an email")
	}
}
