package main

import (
	auth "github.com/Pdv2323/Login-Auth/Auth"
	"github.com/Pdv2323/Login-Auth/login"
	"github.com/Pdv2323/Login-Auth/signin"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// r.GET("/hello", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Welcome to Login Authentication!!")
	// })
	// r.POST("/signin", signin.CreateUser)
	// r.GET("/user")
	// r.GET("/user/{id}")
	// r.DELETE("/user/{id}")

	r.POST("/login", login.UserLogin)
	r.POST("/signup", signin.UserSignUp)
	r.Use(auth.Authz())
	r.GET("/data1", login.GetAll)

	r.Run(":8888")

}
