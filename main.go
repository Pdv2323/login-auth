package main

import (
	"log"

	"github.com/Pdv2323/login-auth/auth"
	database "github.com/Pdv2323/login-auth/db"
	"github.com/Pdv2323/login-auth/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.POST("/login", pkg.UserLogin)
	r.POST("/signup", pkg.UserSignUp)
	r.POST("/otp", pkg.GenerateAndSendOTP)
	r.POST("/changepass", pkg.ChangePasswordWithOTP)
	r.POST("/forgotpass", pkg.ForgetPass)
	r.POST("/resetpass", pkg.ResetPass)
	r.Use(auth.Authz()).GET("/data1", pkg.GetAll)

	r.Run(":1234")

}
