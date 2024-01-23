package pkg

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pdv2323/login-auth/auth"
	"github.com/Pdv2323/login-auth/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"primaryKey"`
	Password string `json:"-" gorm:"not null"`
	OTP      string `json:"otp"`
}

var db *gorm.DB

func UserSignUp(c *gin.Context) {
	var u User

	// err := database.ConnectDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		return
	}

	var existingUser models.User

	result := db.Where("email = ?", u.Email).First(&existingUser)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": true, "message": "User already exists"})
		return
	}

	result = db.Create(&u)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to create user"})
	}

	JwtWrapper1 := auth.JwtWrapper{
		SecretKey:       "esabrfbafbaebhg2425942942",
		Issuer:          "admin",
		ExpirationHours: 48,
	}

	SignedToken, err := JwtWrapper1.GenerateToken(u.Email)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "successfully signed up", "token": SignedToken})
}
