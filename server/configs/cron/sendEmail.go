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
	TeamId   string
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

	if props.TeamId != "" {
		frontendUrl = fmt.Sprintf("%s/dashboard/teams/%s/tasks/%s", os.Getenv("FRONTEND_URL"), props.TeamId, props.TaskId)
	}

	return fmt.Sprintf(configs.DUE_DATE_TEMPLATE, props.Username, props.TaskName, props.DueDate.Local().Format(time.RFC822), frontendUrl)
}

func SendEmailToGmail(to, subject, body string) error {

	emailUrl := os.Getenv("EMAIL_API_URL")

	payload, _ := json.Marshal(EmailRequest{
		Subject: subject, From: "Zendo", To: to, HTML: body,
	})

	req, err := http.NewRequest("POST", emailUrl+"/api/send-email", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", os.Getenv("EMAIL_API_TOKEN"))

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
