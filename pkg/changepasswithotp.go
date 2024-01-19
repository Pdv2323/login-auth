package pkg

import (
	"net/http"

	"github.com/Pdv2323/login-auth/models"
	"github.com/gin-gonic/gin"
)

func ChangePasswordWithOTP(c *gin.Context) {
	var input struct {
		Email       string `json:"email"`
		OTP         string `json:"otp"`
		NewPassword string `json:"new_password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid request payload"})
		return
	}

	var user models.User
	result := db.Where("email = ? AND otp = ?", input.Email, input.OTP).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Invalid OTP"})
		return
	}

	// If OTP is valid, update the password in the database
	user.Password = input.NewPassword
	user.OTP = ""
	// user.OTPExpiry = time.Time{}
	db.Save(&user)

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Password changed successfully"})
}
