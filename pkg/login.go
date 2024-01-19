package pkg

import (
	"log"
	"net/http"

	"github.com/Pdv2323/login-auth/models"
	"github.com/gin-gonic/gin"
)

// var Users map[string]models.User
// var Users = signin.Users

func UserLogin(c *gin.Context) {
	var u models.User

	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		return
	}

	var userfromDB models.User

	result := db.Where("email = ?", u.Email).First(&userfromDB)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "User not found please signup"})
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "User found", "data": userfromDB})
}
