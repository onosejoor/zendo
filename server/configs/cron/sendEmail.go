package cron

import (
	"fmt"
	"main/configs"
	"os"
	"time"

	"gopkg.in/mail.v2"
)

type EmailProps struct {
	Username string
	DueDate  time.Time
	TaskId   string
	TaskName string
}

const (
	SMTPServer = "smtp.gmail.com"
	SMTPPort   = 587
)

func GenerateHtmlTemplate(props EmailProps) string {

	frontendUrl := fmt.Sprintf("%s/dashboard/tasks/%s", os.Getenv("FRONTEND_URL"), props.TaskId)

	return fmt.Sprintf(configs.DUE_DATE_TEMPLATE, props.Username, props.TaskName, props.DueDate.Format(time.RFC822), frontendUrl)
}

func SendEmailToGmail(to, subject, body string) error {
	email := os.Getenv("EMAIL")
	app_password := os.Getenv("APP_PASSWORD")

	mailData := mail.NewMessage()
	mailData.SetHeader("From", fmt.Sprintf("Zendo <%s>", email))
	mailData.SetHeader("To", to)
	mailData.SetHeader("Subject", subject)
	mailData.SetBody("text/html", body)

	// Use 587 for STARTTLS, or 465 for SSL/TLS
	d := mail.NewDialer(SMTPServer, SMTPPort, email, app_password)

	d.StartTLSPolicy = mail.MandatoryStartTLS

	if err := d.DialAndSend(mailData); err != nil {
		return err
	}
	return nil
}
