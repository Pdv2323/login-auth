// package otp

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// func main() {
// 	var userEmail string
// 	fmt.Print("Enter your email: ")
// 	fmt.Scanln(&userEmail)

// 	o := otp.GenerateOtp()

// 	err := otp.SendEmail(userEmail, o)
// 	if err != nil {
// 		log.Fatalf("Error sending email to %s.", userEmail)
// 	}

// 	fmt.Println("OTP sent successfully! Check your email.")

// 	time.Sleep(1000)

// 	var NewOtp int
// 	fmt.Print("Enter the Otp you received : ")
// 	fmt.Scan(&NewOtp)

// 	v := otp.VerifyOtp(o, NewOtp)
// 	fmt.Println(v)

// }
