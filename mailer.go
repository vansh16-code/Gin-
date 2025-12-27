package main

import (
	"net/smtp"
	"os"
)

func SendEmail(to string , subject string , body string)error{

	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")


	addr := host + ":" + port

	msg := []byte(
		"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n",
	)

	auth := smtp.PlainAuth("", user, pass, host)

	return smtp.SendMail(addr, auth, user, []string{to}, msg)




}