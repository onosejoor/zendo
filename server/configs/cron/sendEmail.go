package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/configs"
	"net/http"
	"os"
	"time"
)

type EmailProps struct {
	Username string
	DueDate  time.Time
	TaskId   string
	TaskName string
}

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	HTML    string `json:"html,omitempty"`
	From    string `json:"from"`
}

const (
	SMTPServer = "smtp.gmail.com"
	SMTPPort   = 465
)

func GenerateHtmlTemplate(props EmailProps) string {

	frontendUrl := fmt.Sprintf("%s/dashboard/tasks/%s", os.Getenv("FRONTEND_URL"), props.TaskId)

	return fmt.Sprintf(configs.DUE_DATE_TEMPLATE, props.Username, props.TaskName, props.DueDate.Format(time.RFC822), frontendUrl)
}

func SendEmailToGmail(to, subject, body string) error {

	emailUrl := os.Getenv("EMAIL_API_URL")

	payload, _ := json.Marshal(EmailRequest{
		Subject: subject, From: "Zendo", To: to, HTML: body,
	})

	res, err := http.Post(emailUrl+"/api/send-email", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
