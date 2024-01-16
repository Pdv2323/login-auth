package forgotpass

import (
	"net/http"

	"github.com/Pdv2323/Login-Auth/models"
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

	var user models.User

	result := db.Where("email = ? AND otp = ?", input.Email, input.OTP).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Invalid OTP"})
		return
	}

	user.Password = input.NewPassword
	user.OTP = ""
	db.Save(&user)

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Password reset successfully"})
}
