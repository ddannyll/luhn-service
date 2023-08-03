package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ValidateResponse struct {
	Valid bool `json:"valid"`
}


func validateHandler(c *gin.Context) {
	numberAsString := c.Query("number")
	if _, err := strconv.Atoi(numberAsString); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid number supplied"))
	}
	var sum uint
	for i := 0; i < len(numberAsString) - 1; i++ {
		digitAsNum := numberAsString[i] - '0'
		sum += uint(digitAsNum)
	}
	c.JSON(http.StatusOK, gin.H{
		"valid": sum % 10 == uint(numberAsString[len(numberAsString) - 1] - '0'),
	})
}