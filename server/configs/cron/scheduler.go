package cron

import (
	"context"
	"fmt"
	"log"
	"main/db"
	"time"

	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var scheduler *gocron.Scheduler

func InitializeGoCron() {
	if scheduler != nil {
		return
	}

	scheduler = gocron.NewScheduler(time.Local)
	scheduler.StartAsync()
	log.Println("Cron scheduler active")
}

func ScheduleReminderJob(taskID primitive.ObjectID, dueDate time.Time, userID primitive.ObjectID, userEmail string) (*gocron.Job, error) {
	reminderTime := calculateReminderTime(dueDate)

	if reminderTime.Before(time.Now()) {
		log.Printf("Reminder for task %s is in the past or too soon. Handling immediately. \n", taskID)
		return nil, fmt.Errorf("due date is too soon to schedule a future reminder")
	}

	job, err := scheduler.
		Every(1).Minute().
		At(reminderTime.Format("15:04:05")).
		Tag(taskID.Hex()).
		Do(sendEmailReminder, taskID, userID, userEmail, dueDate)

	if err != nil {
		log.Printf("Error scheduling reminder for task %s: %v", taskID, err)
		return nil, err
	}

	log.Printf("Reminder scheduled for task %s at %s \n", taskID, reminderTime.Format(time.RFC3339))
	return job, nil
}

func calculateReminderTime(dueDate time.Time) time.Time {
	diff := time.Until(dueDate)

	if diff <= 5*time.Minute {
		// If less than or equal to 5 minutes, remind at 5 minutes before
		return dueDate.Add(-5 * time.Minute)
	} else {
		// Otherwise, remind at 1 hour before
		return dueDate.Add(-1 * time.Hour)
	}
}

func sendEmailReminder(taskID primitive.ObjectID, userID primitive.ObjectID, userEmail string, dueDate time.Time) {
	log.Printf("Sending reminder for task ID: %s to user %s (%s). Due: %s", taskID, userID, userEmail, dueDate.Format(time.RFC822))

	// send mail here ....

	// if no error then

	client := db.GetClient().Collection("tasks")

	filter := bson.M{
		"_id":    taskID,
		"userId": userID,
	}

	_, err := client.UpdateOne(context.TODO(), filter, bson.M{
		"$set": bson.M{
			"reminder_sent": true,
		}})
	if err != nil {
		log.Println("Error updating task")
	}

	// TODO: Implement your actual email sending logic here
	// Example: notifications.SendEmail(userEmail, "Task Reminder", fmt.Sprintf("Your task '%s' is due soon!", taskID))

	// After sending, you might want to mark the reminder as sent in your DB
	// to prevent duplicate reminders if your scheduler somehow re-runs.
	// As discussed, this isn't strictly necessary if your logic only schedules one-offs
	// and doesn't reschedule past events, but it's good for robustness.
	// For instance, you could update the task: task.ReminderSent = true
}
