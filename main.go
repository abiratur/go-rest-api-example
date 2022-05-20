package main

import (
	"bufio"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type validatePasswordRequest struct {
	Password string `json:"password"`
}

type validatePasswordResponse struct {
	IsValid bool `json:"isValid"`
}

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

func validatePassword(c *gin.Context) {

	var request validatePasswordRequest

	if err := c.BindJSON(&request); err != nil {
		println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response := validatePasswordResponse{IsValid: true}

	if passwords[request.Password] {
		response.IsValid = false
	}

	c.IndentedJSON(http.StatusOK, response)
}
