package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func readPasswords(passwordsFilePath string) (map[string]bool, error) {
	file, err := os.Open(passwordsFilePath)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("cannot open file")
	}
	defer file.Close()

	passwords := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passwords[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, errors.New("scanner error")
	}

	return passwords, nil
}
