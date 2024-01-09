package signin

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/Pdv2323/Login-Auth/JWT"
	"github.com/Pdv2323/Login-Auth/models"
	"github.com/gin-gonic/gin"
)

var Users map[string]models.User

func UserSignUp(c *gin.Context) {
	var u models.User

	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		return
	}

	_, ok := Users[u.Email]
	if ok {
		c.JSON(http.StatusConflict, gin.H{"message": "You are already registered please sign up"})
		return
	}

	Users = map[string]models.User{
		u.Email: u,
	}

	JwtWrapper1 := jwt.JwtWrapper{
		SecretKey:       "esabrfbafbaebhg2425942942",
		Issuer:          "admin",
		ExpirationHours: 48,
	}

	SignedToken, jwtErr := JwtWrapper1.GenerateToken(u.Email)
	if jwtErr != nil {
		fmt.Println(jwtErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "successfully signed up", "token": SignedToken})
}
