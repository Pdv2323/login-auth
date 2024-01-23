package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResetPass(c *gin.Context) {
	var input struct {
		Email       string `json:"email"`
		OTP         string `json:"otp"`
		NewPassword string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
		return
	}

	var u User

	result := db.Where("email = ? AND otp = ?", input.Email, input.OTP).First(&u)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Invalid OTP"})
		return
	}

	u.Password = input.NewPassword
	u.OTP = ""
	db.Save(&u)

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Password reset successfully"})
}
