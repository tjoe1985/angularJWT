package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello there",
		})
	}
}
