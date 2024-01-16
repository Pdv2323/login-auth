package onetimepass

import (
	"math/rand"
	"time"
)

// func GenerateOtp() string {
// 	rand.Seed(time.Now().UnixNano())
// 	r_value := rand.Intn(900000) + 100000
// 	s := string(r_value)
// 	return s
// }

func GenerateOtp() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000) + 100000
}
