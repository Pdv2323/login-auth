package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type NewUser struct {
	gorm.Model
	Id         int    `gorm:"primary_key" json:"id"`
	First_Name string `gorm:"not null" json:"first_name"`
	Last_Name  string `gorm:"not null" json:"last_name"`
	Email_Id   string `gorm:"not null" json:"email_id"`
	Password   string `gorm:"not null" json:"-"`
	Otp        int    `gorm:"not null" json:"otp"`
}

type UserLogin struct {
	gorm.Model
	Email_Id string ``
}

func (u *NewUser) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *NewUser) CheckPasswordHash(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

//New Structs

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}
