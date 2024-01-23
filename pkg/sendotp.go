package pkg

import (
	"log"
	"net/http"

	onetimepass "github.com/Pdv2323/login-auth/otp"
	"github.com/gin-gonic/gin"
)

func GenerateAndSendOTP(c *gin.Context) {
	//email string
	var input struct {
		Email string `json:"email"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid Request"})
		return
	}
	otp := onetimepass.GenerateOtp()

	// Save OTP and its expiry time in the database
	var user User
	result := db.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		log.Println("User not found")
		return
	}

	user.OTP = otp
	// user.OTPExpiry = time.Now().Add(5 * time.Minute) // Set OTP expiry time to 5 minutes
	db.Save(&user)

	// Send OTP to the user
	onetimepass.SendEmail(input.Email, otp)
}
