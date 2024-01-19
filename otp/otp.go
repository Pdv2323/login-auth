package onetimepass

import (
	"math/rand"
)

// func GenerateOtp() string {
// 	rand.Seed(time.Now().UnixNano())
// 	r_value := rand.Intn(900000) + 100000
// 	s := string(r_value)
// 	return s
// }

// func GenerateOtp() int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(900000) + 100000
// }

func GenerateOtp() string {
	// Generate a random 6-digit OTP
	const otpLength = 6
	const digits = "0123456789"

	otp := make([]byte, otpLength)
	for i := range otp {
		otp[i] = digits[rand.Intn(len(digits))]
	}
	return string(otp)
}
