// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	auth "github.com/Pdv2323/Login-Auth/Auth"
// 	jwt "github.com/Pdv2323/Login-Auth/JWT"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// const (
// 	dbHost     = "localhost"
// 	dbPort     = 5432
// 	dbName     = "postgres"
// 	dbUser     = "postgres"
// 	dbPassword = "123"
// )

// var Users map[string]User

// func UserSignUp(c *gin.Context) {
// 	var u User
// 	ConnectDB()

// 	err := c.BindJSON(&u)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	_, ok := Users[u.Email]
// 	if ok {
// 		c.JSON(http.StatusConflict, gin.H{"message": "You are already registered please login."})
// 		return
// 	}

// 	Users = map[string]User{
// 		u.Email: u,
// 	}

// 	JwtWrapper1 := jwt.JwtWrapper{
// 		SecretKey:       "esabrfbafbaebhg2425942942",
// 		Issuer:          "admin",
// 		ExpirationHours: 48,
// 	}

// 	SignedToken, err := JwtWrapper1.GenerateToken(u.Email)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"error": false, "message": "successfully signed up", "token": SignedToken})
// }

// func UserLogin(c *gin.Context) {
// 	var u User
// 	ConnectDB()

// 	err := c.BindJSON(&u)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	val, ok := Users[u.Email]
// 	if !ok {
// 		c.JSON(http.StatusConflict, gin.H{"message": "Please signup User does not exist."})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Successfully logged in", "data": val})
// }

// func GetAll(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"data": Users})
// }

// func ConnectDB() (db *gorm.DB, err error) {
// 	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
// 		dbHost, dbPort, dbName, dbUser, dbPassword)
// 	db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func main() {

// 	ConnectDB()

// 	r := gin.Default()

// 	// r.GET("/hello", func(c *gin.Context) {
// 	// 	c.String(http.StatusOK, "Welcome to Login Authentication!!")
// 	// })
// 	// r.POST("/signin", signin.CreateUser)
// 	// r.GET("/user")
// 	// r.GET("/user/{id}")
// 	// r.DELETE("/user/{id}")

// 	r.POST("/login", UserLogin)
// 	r.POST("/signup", UserSignUp)
// 	r.Use(auth.Authz())
// 	r.GET("/data1", GetAll)

// 	r.Run(":8000")

// }
package main

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/Pdv2323/Login-Auth/Auth"
	jwt "github.com/Pdv2323/Login-Auth/JWT"
	onetimepass "github.com/Pdv2323/Login-Auth/otp"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	OTP      int    `json:"otp" gorm:"int"`
}

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbName     = "postgres"
	dbUser     = "postgres"
	dbPassword = "123"
)

var db *gorm.DB

// var Users map[string]User
var Users User

func UserSignUp(c *gin.Context) {
	var u User

	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		return
	}

	var existingUser User
	result := db.Where("email = ?", u.Email).First(&existingUser)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": true, "message": "User already exists"})
		return
	}

	result = db.Create(&u)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to create user"})
	}

	// c.JSON(http.StatusOK, gin.H{"error": false, "message": "User created successfully."})
	// 	Users = map[string]User{
	// 		u.Email: u,
	// 	}

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
	// ConnectDB()

	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		return
	}

	var userfromDB User
	result := db.Where("email = ?", u.Email).First(&userfromDB)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "User not found please signup"})
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "User found", "data": userfromDB})
}

func GetAll(c *gin.Context) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func ConnectDB() error {
	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPassword)

	var err error
	db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Users)
	if err != nil {
		return err
	}
	return nil
}

func ForgetPass(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid Request"})
		return
	}

	var u User

	result := db.Where("email = ?", input.Email).First(&u)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "User not found or Invalid email"})
	}

	otp := onetimepass.GenerateOtp()
	err := onetimepass.SendEmail(input.Email, otp)
	if err != nil {
		log.Fatalf("Error sending email to %s.", input.Email)
	}

	fmt.Println("OTP sent successfully! Check your email.")
	u.OTP = otp
	db.Save(&u)

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Reset token generated successfully"})

}
func ResetPass(c *gin.Context) {
	var input struct {
		Email       string `json:"email"`
		OTP         int    `json:"otp"`
		NewPassword string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
		return
	}

	var u User

	result := db.Where("email = ? AND otp = ?", input.Email, input.OTP).First(&u)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Invalid OTP"})
		return
	}

	u.Password = input.NewPassword
	u.OTP = input.OTP
	db.Save(&u)

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Password reset successfully"})
}

func main() {

	err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

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
	r.POST("/forgotpass", ForgetPass)
	r.POST("/resetpass", ResetPass)

	r.Run(":8000")

}
