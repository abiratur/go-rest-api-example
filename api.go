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
