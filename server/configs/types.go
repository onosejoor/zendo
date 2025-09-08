package configs

type GooglePayload struct {
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

const EMAIL_TEMPLATE = `
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
            background-color: #3b82f6;
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
            background-color: #2563eb;
        }

        .greeting {
            margin-bottom: 24px;
        }

        .message-text {
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

        <p class="message-text">You're almost there! Click the button below to verify your email and sign in securely:</p>

        <div class="action-section">
            <a href="%s" class="button">Verify Email & Sign In</a>
        </div>

        <p>This magic link will expire in 15 minutes. If you didn’t request it, you can ignore this email.</p>


        <div class="signature">
            <p>Thanks for using Zendo!<br><br>
                <span class="team-name">Zendo team</span>
            </p>
        </div>
    </div>
</body>

</html>
    `

const DUE_DATE_TEMPLATE = `
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

        .greeting {
            margin-bottom: 24px;
        }

        .reminder-text {
            margin-bottom: 24px;
        }

        .task-info {
            background-color: #fef3c7;
            padding: 20px;
            border-radius: 6px;
            margin: 24px 0;
            border-left: 4px solid #f59e0b;
        }

        .task-title {
            font-size: 18px;
            font-weight: 600;
            color: #1d2025ff;
            margin: 0 0 8px 0;
        }

        .due-date {
            font-style: italic;
            font-weight: 500;
            color: #ef4343ff;
            margin: 0;
        }

        .action-section {
            margin: 32px 0;
            text-align: center;
        }

        .signature {
            margin-top: 32px;
            padding-top: 16px;
        }

        .team-name {
            font-weight: 600;
        }

        .secondary-text {
            color: #64748b;
            font-size: 14px;
            margin-top: 16px;
        }

        @media (max-width: 600px) {
            .email-container {
                padding: 20px;
                margin: 10px;
            }
            
            body {
                padding: 10px;
            }
            
            .task-info {
                padding: 16px;
            }
        }
    </style>
</head>

<body>
    <div class="email-container">
        <p class="greeting">Hi %s,</p>

        <p class="reminder-text">⏰ This is a quick reminder that one of your tasks is coming due soon:</p>

        <div class="task-info">
            <p class="task-title">Task: %s</p>
            <p class="due-date">Due Date: %s</p>
        </div>

        <p class="reminder-text">Please make sure to wrap it up before the deadline.</p>

        <div class="action-section">
            <a href="%s" class="button">View Task Details</a>
        </div>

        <p class="secondary-text">Need help or have questions? Just reply to this email and we'll be happy to assist.</p>

        <div class="signature">
            <p>Let us know if you need any help.<br><br>
                Best regards,<br>
                <span class="team-name">Zendo team</span>
            </p>
        </div>
    </div>
</body>

</html>
    `

const TEAM_INVITATION_EMAIL_TEMPLATE = `
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
            background-color: #3b82f6;
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
            background-color: #2563eb;
        }

        .greeting {
            margin-bottom: 24px;
        }

        .message-text {
            margin-bottom: 24px;
        }

        .team-info {
            background-color: #f8fafc;
            padding: 20px;
            border-radius: 6px;
            margin: 24px 0;
            border-left: 4px solid #3b82f6;
        }

        .team-name-highlight {
            font-weight: 600;
            color: #1e40af;
        }

        .inviter-name {
            font-weight: 500;
        }

        .action-section {
            margin: 32px 0;
            text-align: center;
        }

        .signature {
            margin-top: 32px;
            padding-top: 16px;
        }

        .team-name {
            font-weight: 600;
        }

        .secondary-text {
            color: #64748b;
            font-size: 14px;
            margin-top: 16px;
        }

        @media (max-width: 600px) {
            .email-container {
                padding: 20px;
                margin: 10px;
            }
            
            body {
                padding: 10px;
            }
            
            .team-info {
                padding: 16px;
            }
        }
    </style>
</head>

<body>
    <div class="email-container">
        <p class="greeting">Hi There 👋,</p>

        <p class="message-text">Great news! <span class="inviter-name">%s</span> has invited you to collaborate on their team in Zendo.</p>

        <div class="team-info">
            <p><strong>Team:</strong> <span class="team-name-highlight">%s</span></p>
            <p><strong>Role:</strong> %s</p>
            <p><strong>Invited by:</strong> %s</p>
        </div>

        <p class="message-text">Join the team to start collaborating, sharing projects, and working together seamlessly.</p>

        <div class="action-section">
            <a href="%s" class="button">Accept Invitation</a>
        </div>

        <p class="secondary-text">This invitation will expire in 7 days. If you don't want to join this team, you can safely ignore this email.</p>

        <div class="signature">
            <p>Welcome to better collaboration!<br><br>
                <span class="team-name">Zendo team</span>
            </p>
        </div>
    </div>
</body>

</html>
    `
