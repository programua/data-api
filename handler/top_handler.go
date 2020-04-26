package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TopHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "top ok gin!"})
}