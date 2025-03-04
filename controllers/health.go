package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Receive a pointer to a Context
// gin.Context contains all info abt request and allows to return a response
func ControllerHealthCheck(c *gin.Context) {
	// IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body.
	// Includes de status onf http and the struct
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successful Health Check."})
}
