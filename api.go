package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PasswordsController struct {
	// go doesn't have a 'hashset' type :/
	// so we use a map instead: https://stackoverflow.com/questions/34018908/golang-why-dont-we-have-a-set-datastructure
	Passwords map[string]bool
}

type validatePasswordRequest struct {
	Password string `json:"password"`
}

type validatePasswordResponse struct {
	IsValid bool `json:"isValid"`
}

func (controller *PasswordsController) validatePassword(c *gin.Context) {

	var request validatePasswordRequest

	if err := c.BindJSON(&request); err != nil {
		println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response := validatePasswordResponse{IsValid: true}

	if controller.Passwords[request.Password] {
		response.IsValid = false
	}

	c.IndentedJSON(http.StatusOK, response)
}
