package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}
