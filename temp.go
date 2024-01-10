package main

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/Pdv2323/Login-Auth/Auth"
	jwt "github.com/Pdv2323/Login-Auth/JWT"
	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Users map[string]User

func UserSignUp(c *gin.Context) {
	var u User

	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		return
	}

	_, ok := Users[u.Email]
	if ok {
		c.JSON(http.StatusConflict, gin.H{"message": "You are already registered please login."})
		return
	}

	Users = map[string]User{
		u.Email: u,
	}

	JwtWrapper1 := jwt.JwtWrapper{
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

func UserLogin(c *gin.Context) {
	var u User

	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		return
	}

	val, ok := Users[u.Email]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"message": "Please signup User does not exist."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Successfully logged in", "data": val})
}

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": Users})
}

func main() {
	r := gin.Default()

	// r.GET("/hello", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Welcome to Login Authentication!!")
	// })
	// r.POST("/signin", signin.CreateUser)
	// r.GET("/user")
	// r.GET("/user/{id}")
	// r.DELETE("/user/{id}")

	r.POST("/login", UserLogin)
	r.POST("/signup", UserSignUp)
	r.Use(auth.Authz())
	r.GET("/data1", GetAll)

	r.Run(":8000")

}
