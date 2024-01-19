package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"primaryKey"`
	Password string `json:"-" gorm:"not null"`
	OTP      string `json:"otp"`
}
