package database

import (
	"fmt"
	"net/http"

	jwt "github.com/Pdv2323/Login-Auth/JWT"
	"github.com/Pdv2323/Login-Auth/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SignupToDatabase(c *gin.Context) {
	ConnectDB()
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error invalid Request"})
		return
	}

	JwtWrapper1 := jwt.JwtWrapper{
		SecretKey:       "esabrfbafbaebhg2425942942",
		Issuer:          "admin",
		ExpirationHours: 48,
	}

	SignedToken, err := JwtWrapper1.GenerateToken(user.Email)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username alreayd exists."})
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "successfully signed upto database", "token": SignedToken})

}

func LoginUsingDatabase(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// val, ok := user[
	// if !ok {
	// 	c.JSON(http.StatusConflict, gin.H{"message": "Please signup User does not exist."})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"error": false, "message": "Successfully logged in", "data": val})

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Successfully logged in", "data": user})

	// // Generate JWT token
	// tokenString, err := generateToken(int(user.ID), user.Username)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
