package cron

import (
	"fmt"
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

	return fmt.Sprintf(`
    <!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Geist:wght@100..900&display=swap" rel="stylesheet">
    <style>
        * {
         font-family: "Geist", sans-serif;
        }

        body {
            padding: 20px;
            color: #282c34ff;
            background-color: #f9f9f9ff;
            margin: 0;
            line-height: 1.6;
        }

        .email-container {
            background-color: #ffffffff;
            padding: 32px;
            border-radius: 8px;
            max-width: 600px;
            margin: auto;
            box-shadow: 0 10px 25px -5px #0000001a;
        }

        .button {
            background-color: #16a249ff;
            color: white;
            padding: 12px 24px;
            text-align: center;
            text-decoration: none;
            display: inline-block;
            border-radius: 6px;
            font-weight: 600;
            font-size: 14px;
            transition: background-color 0.2s ease;
            border: none;
            cursor: pointer;
        }

        .button:hover {
            background-color: #149041ff;
        }

        .task-title {
            font-size: 18px;
            font-weight: 600;
            color: #1d2025ff;
            margin: 16px 0 8px 0;
        }

        .due-date {
            font-style: italic;
            font-weight: 500;
            color: #ef4343ff;
        }

        .greeting {
            margin-bottom: 24px;
        }

        .reminder-text {
            margin-bottom: 24px;
        }

        .task-details {
            margin-bottom: 24px;
        }

        .action-section {
            margin: 32px 0;
        }

        .signature {
            margin-top: 32px;
            padding-top: 16px;
        }

        .team-name {
            font-weight: 600;
        }

        @media (max-width: 600px) {
            .email-container {
                padding: 20px;
                margin: 10px;
            }
            
            body {
                padding: 10px;
            }
        }
    </style>
</head>

<body>
    <div class="email-container">
        <p class="greeting">Hi %s,</p>

        <p class="reminder-text">⏰ This is a quick reminder that one of your tasks is coming due soon:</p>

        <div class="task-details">
            <p class="task-title">Task: %s</p>
            <p>Due Date: <span class="due-date">%s</span></p>
        </div>

        <p>Please make sure to wrap it up before the deadline.</p>

        <div class="action-section">
            <a href="%s" class="button">View Task Details</a>
        </div>

        <div class="signature">
            <p>Let us know if you need any help.<br><br>
                Best regards,<br>
                <span class="team-name">Zendo team</span>
            </p>
        </div>
    </div>
</body>

</html>
    `, props.Username, props.TaskName, props.DueDate.Format(time.RFC822), frontendUrl)
}

func SendEmailToGmail(to, subject, body string) error {
	email := os.Getenv("EMAIL")
	app_password := os.Getenv("APP_PASSWORD")

	mailData := mail.NewMessage()
	mailData.SetHeader("From", email)
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
