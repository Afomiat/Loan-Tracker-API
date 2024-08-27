package infrastructure

import (
    "gopkg.in/mail.v2"
)

type EmailService struct {
    SMTPHost     string
    SMTPPort     string
    SMTPUsername string
    SMTPPassword string
}

func (es *EmailService) SendVerificationEmail(to, token string) error {
    m := mail.NewMessage()
    m.SetHeader("From", es.SMTPUsername)
    m.SetHeader("To", to)
    m.SetHeader("Subject", "Email Verification")
    m.SetBody("text/plain", "Please verify your email using this link: /users/verify-email?token="+token+"&email="+to)

    d := mail.NewDialer(es.SMTPHost, 587, es.SMTPUsername, es.SMTPPassword)
    return d.DialAndSend(m)
}
