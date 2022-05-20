package main

import (
	"net/http"

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
	router := gin.Default()
	router.POST("/v1/validatePassword", validatePassword)

	router.Run("localhost:8080")
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
