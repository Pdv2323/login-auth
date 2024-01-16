package onetimepass

import (
	"fmt"
	"log"
	"time"
)

func TestRun() {
	var userEmail string
	fmt.Print("Enter your email: ")
	fmt.Scanln(&userEmail)

	o := GenerateOtp()

	err := SendEmail(userEmail, o)
	if err != nil {
		log.Fatalf("Error sending email to %s.", userEmail)
	}

	fmt.Println("OTP sent successfully! Check your email.")

	time.Sleep(1000)

	var NewOtp int
	fmt.Print("Enter the Otp you received : ")
	fmt.Scan(&NewOtp)

	v := VerifyOtp(o, NewOtp)
	fmt.Println(v)

}
