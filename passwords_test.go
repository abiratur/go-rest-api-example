package main

import (
	"testing"
)

func PasswordsFileShouldBeRead(t *testing.T) {

	_, err := readPasswords("passwords.txt")
	if err != nil {
		t.Fatalf(`got error from readPasswords: %v`, err)
	}
}
