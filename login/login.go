package login

import (
	"log"
	"net/http"

	"github.com/Pdv2323/Login-Auth/models"
	"github.com/gin-gonic/gin"
)

var Users map[string]models.User

// var Users = signin.Users

func UserLogin(c *gin.Context) {
	var u models.User

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
