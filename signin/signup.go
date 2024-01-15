package signin

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/Pdv2323/Login-Auth/JWT"
	"github.com/Pdv2323/Login-Auth/models"
	"github.com/gin-gonic/gin"
)

var Users map[string]models.User
var db *sql.DB

// var Users models.User

func UserSignUp(c *gin.Context) {
	var u models.User

	// db, err := db.ConnectDB()
	// if err != nil {
	// 	log.Println(err)
	// }
	// db.ConnectDB()

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

	Users = map[string]models.User{
		u.Email: u,
		// u.Password: u,
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
