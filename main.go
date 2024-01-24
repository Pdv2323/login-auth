package main

import (
	"log"

	"github.com/Pdv2323/login-auth/auth"
	database "github.com/Pdv2323/login-auth/db"
	"github.com/Pdv2323/login-auth/pkg"
	"github.com/gin-gonic/gin"
)

// var db *gorm.DB

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		// Check if db is not nil before closing
		if db != nil {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}
	}()

	// defer db.Close()

	r := gin.Default()

	// r.Use(func(c *gin.Context) {
	// 	c.Set("db", db)
	// 	c.Next()
	// })

	// r.Use(gin.Recovery())

	var srv = pkg.Server{
		Db: db,
	}

	r.POST("/login", pkg.UserLogin)
	r.POST("/signup", srv.UserSignUp)
	r.POST("/otp", pkg.GenerateAndSendOTP)
	r.POST("/changepass", pkg.ChangePasswordWithOTP)
	r.POST("/forgotpass", pkg.ForgetPass)
	r.POST("/resetpass", pkg.ResetPass)
	r.Use(auth.Authz()).GET("/data1", pkg.GetAll)

	r.Run(":1234")

}
