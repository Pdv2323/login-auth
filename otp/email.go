package onetimepass

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to string, otp string) error {
	from := "parthdv2111@gmail.com"
	password := "dwvntkfkeqnejzco"
	SmtpServer := "smtp.gmail.com"
	SmtpPort := "587"

	msg := fmt.Sprintf("Subject : OTP Verification\n\n Your OTP is : %s", otp)

	auth := smtp.PlainAuth("", from, password, SmtpServer)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", SmtpServer, SmtpPort),
		auth,
		from,
		[]string{to},
		[]byte(msg),
	)

	return err
}
