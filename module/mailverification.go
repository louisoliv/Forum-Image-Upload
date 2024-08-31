package module

import (
	"fmt"
	"net/smtp"
)

func MailVerification(email, token string) error {
	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUsername := "jeremy.cailly76@gmail.com"
	smtpPassword := "tauaytjuyuthsjox"

	// Sender and recipient email addresses
	sender := "jeremy.cailly76@gmail.com"
	recipient := email

	// Craft email template with verification link
	verificationLink := fmt.Sprintf("http://localhost:8000/verify?token=%s", token)
	emailBody := fmt.Sprintf("Click the following link to verify your email: %s", verificationLink)

	// Email content
	subject := "Email Verification"
	message := "Subject: " + subject + "\r\n" +
		"\r\n" +
		emailBody

	// SMTP authentication
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Send email using SMTP
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, []string{recipient}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
