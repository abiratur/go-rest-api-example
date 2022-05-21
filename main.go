package main

import (
	"github.com/gin-gonic/gin"
)

// go doesn't have a 'hashset' type :/
// so we use a map instead: https://stackoverflow.com/questions/34018908/golang-why-dont-we-have-a-set-datastructure
var passwords = map[string]bool{
	"123456":    true,
	"password!": true,
}

func main() {

	var err error
	passwords, err = readPasswords("passwords.txt")
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.POST("/v1/validatePassword", validatePassword)

	router.Run("0.0.0.0:8080")
}
