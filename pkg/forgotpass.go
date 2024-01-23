package pkg

import (
	"log"
	"net/http"

	onetimepass "github.com/Pdv2323/login-auth/otp"
	"github.com/gin-gonic/gin"
)

func ForgetPass(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid Request"})
		return
	}

	var u User

	result := db.Where("email = ?", input.Email).First(&u)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "User not found or Invalid email"})
	}

	otp := onetimepass.GenerateOtp()

	u.OTP = otp
	db.Save(&u)

	err := onetimepass.SendEmail(input.Email, otp)
	if err != nil {
		log.Fatalf("Error sending email to %s.", input.Email)
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Reset token generated successfully"})

}
