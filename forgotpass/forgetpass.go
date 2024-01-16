package forgotpass

import (
	"net/http"

	"github.com/Pdv2323/Login-Auth/models"
	"github.com/Pdv2323/Login-Auth/otp"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func ForgetPass(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid Request"})
		return
	}

	var user models.User

	result := db.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "User not found or Invalid email"})
	}

	otp := otp.GenerateOtp()
	user.OTP = otp
	db.Save(&user)

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Reset token generated successfully"})

}
