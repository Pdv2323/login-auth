package otp

import (
	"math/rand"
	"time"
)

func GenerateOtp() string {
	rand.Seed(time.Now().UnixNano())
	r_value := rand.Intn(900000) + 100000
	return string(r_value)
}
