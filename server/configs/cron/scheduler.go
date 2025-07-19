package cron

import (
	"context"
	"fmt"
	"log"
	"main/db"
	"main/models"
	"time"

	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReminderProps struct {
	TaskID   primitive.ObjectID
	DueDate  time.Time
	UserID   primitive.ObjectID
	TaskName string
	Ctx      context.Context
}

var scheduler *gocron.Scheduler

// this function initializes the cronjobs in the backhround
func InitializeGoCron() {
	if scheduler != nil {
		return
	}

	scheduler = gocron.NewScheduler(time.Local)
	scheduler.StartAsync()
	log.Println("Cron scheduler active")
}

// a method to schedule a reminder
func (params ReminderProps) ScheduleReminderJob() error {
	reminderTime := calculateReminderTime(params.DueDate)

	if reminderTime.Before(time.Now().Local()) {
		log.Printf("Reminder for task %s is in the past or too soon. Handling immediately. \n", params.TaskID)
		return fmt.Errorf("due date is too soon to schedule a future reminder")
	}

	_, err := scheduler.Every(1).Day().
		Tag(params.TaskID.Hex()).
		At(reminderTime).
		Do(params.sendEmailReminder)
	if err != nil {
		log.Printf("Error scheduling reminder for task %s: %v", params.TaskID, err)
		return err
	}

	log.Printf("Reminder scheduled for task %s at %s \n", params.TaskID, reminderTime.Format(time.RFC3339))
	return nil
}
func calculateReminderTime(due time.Time) time.Time {
	diff := time.Until(due)
	switch {
	case diff <= 5*time.Minute:
		adjusted := time.Now().Add(30 * time.Second)
		log.Println("🟡 Due soon (<=5min). Reminder in 30s:", adjusted)
		return adjusted

	case diff <= 1*time.Hour:
		adjusted := due.Add(-diff / 2)
		log.Println("🟠 Due within an hour. Reminder at midpoint:", adjusted)
		return adjusted

	default:
		adjusted := due.Add(-1 * time.Hour)
		log.Println("🟢 Due in > 1hr. Reminder 1hr before:", adjusted)
		return adjusted
	}
}

func (params ReminderProps) sendEmailReminder() {
	client := db.GetClient()
	tasksCollection := client.Collection("tasks")
	usersCollection := client.Collection("users")

	user, err := models.GetUser(params.UserID, usersCollection, params.Ctx)
	if err != nil {
		log.Println("[Reminder] Error getting user data:", err)
		return
	}

	log.Printf("[Reminder] Sending for task ID: %s to user. Due: %s",
		params.TaskID.Hex(), params.DueDate.Format(time.RFC822))

	reminderSent, err := models.GetTaskReminderSent(params.TaskID, tasksCollection, params.Ctx)
	if err != nil {
		log.Println("[Reminder] Error fetching task reminder flag:", err)
		return
	}

	if reminderSent {
		log.Println("[Reminder] Already sent for task:", params.TaskID.Hex())
		return
	}

	htmlTemplate := GenerateHtmlTemplate(EmailProps{
		Username: user.Username,
		TaskId:   params.TaskID.Hex(),
		DueDate:  params.DueDate,
		TaskName: params.TaskName,
	})

	sendErr := SendEmailToGmail(user.Email, "Task Getting Due", htmlTemplate)
	if sendErr != nil {
		log.Printf("[Reminder] Initial email send failed: %v", sendErr)
		time.Sleep(5 * time.Second)

		sendErr = SendEmailToGmail(user.Email, "Task Getting Due", htmlTemplate)
		if sendErr != nil {
			log.Printf("[Reminder] Retry email send failed: %v", sendErr)
			return
		}
	}

	filter := bson.M{
		"_id":    params.TaskID,
		"userId": params.UserID,
	}

	_, err = tasksCollection.UpdateOne(params.Ctx, filter, bson.M{
		"$set": bson.M{
			"reminder_sent": true,
		},
	})
	if err != nil {
		log.Printf("[Reminder] Error updating task %s for user %s: %v", params.TaskID.Hex(), params.UserID.Hex(), err)
	} else {
		log.Printf("[Reminder] Email sent successfully and reminder marked as sent for task %s", params.TaskID.Hex())
	}
}
