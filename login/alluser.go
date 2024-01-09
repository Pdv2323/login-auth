package login

import (
	"net/http"

	"github.com/Pdv2323/Login-Auth/models"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var u models.User
	c.JSON(http.StatusOK, gin.H{"data": u})
}
