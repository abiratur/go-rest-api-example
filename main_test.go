package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {

	_, err := readPasswords("passwords.txt")
	if err != nil {
		t.Fatalf(`go error from readPasswords: %v`, err)
	}
}
