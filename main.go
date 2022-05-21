package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	passwords, err := readPasswords("passwords.txt")
	if err != nil {
		panic(err)
	}

	passwordsController := PasswordsController{Passwords: passwords}

	router := gin.Default()
	router.POST("/v1/validatePassword", passwordsController.validatePassword)

	router.Run("0.0.0.0:8080")
}
