package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Authz() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("Authorization")
		if ClientToken == "" {
			c.JSON(403, "No authorization token provided.")
			c.Abort()
			return
		}

		ExtractedToken := strings.Split(ClientToken, "Bearer")

		if len(ExtractedToken) == 2 {
			ClientToken = strings.TrimSpace(ExtractedToken[1])
		} else {
			c.JSON(400, "Incorrect format of authorization token.")
			c.Abort()
			return
		}

		JwtWrapper1 := JwtWrapper{
			SecretKey: "esabrfbafbaebhg2425942942",
			Issuer:    "admin",
		}

		claims, err := JwtWrapper1.ValidateToken(ClientToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}

}
