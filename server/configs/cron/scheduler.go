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
func InitializeGoCron() *gocron.Scheduler {
	if scheduler != nil {
		return scheduler
	}

	scheduler = gocron.NewScheduler(time.Local)
	scheduler.StartAsync()

	scheduler.Every(10).Minutes().Do(func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[Cron] Panic recovered in SetTasksCron: %v", r)
			}
		}()

		if err := SetTasksCron(context.Background()); err != nil {
			log.Println("[Cron] Error in SetTasksCron:", err)
		}
	})
	log.Println("Cron scheduler active")

	return scheduler
}

func DeleteCronJob(taskId primitive.ObjectID) error {
	jobs, err := scheduler.FindJobsByTag(taskId.Hex())
	if err != nil {
		log.Printf("Error finding cron job for task %s: %v", taskId.Hex(), err)
		return err
	}

	if len(jobs) < 1 {
		log.Printf("No cron job found for task %s, nothing to delete", taskId.Hex())
		return nil
	}

	err = scheduler.RemoveByTag(taskId.Hex())
	if err != nil {
		log.Printf("Error removing cron job for task %s: %v", taskId.Hex(), err)
		return err
	}

	log.Printf("Cron job deleted for task %s", taskId.Hex())
	return nil
}

// a method to schedule a reminder
func (params ReminderProps) ScheduleReminderJob() error {
	reminderTime := calculateReminderTime(params.DueDate)

	if reminderTime.Before(time.Now().Local()) {
		log.Printf("Reminder for task %s is in the past or too soon. Handling immediately. \n", params.TaskID)
		return fmt.Errorf("due date is too soon to schedule a future reminder")
	}

	_, err := scheduler.Every(1).Day().StartAt(reminderTime).LimitRunsTo(1).
		Tag(params.TaskID.Hex()).
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
	case diff <= 10*time.Minute:
		adjusted := time.Now().Add(30 * time.Second)
		log.Println("🟡 Due soon (<=10min). Reminder in 30s:", adjusted)
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
	usersCollection := client.Collection("users")

	user, err := models.GetUser(params.UserID, usersCollection, params.Ctx)
	if err != nil {
		log.Println("[Reminder] Error getting user data:", err)
		return
	}

	reminderDoesNotExists, err := models.GetTaskReminderSent(params.TaskID, client, params.Ctx)
	if err != nil {
		log.Println("[Reminder] Error fetching task reminder flag:", err)
		return
	}

	if reminderDoesNotExists {
		log.Println("[Reminder] Already sent or does not exist for task:", params.TaskID.Hex())
		return
	}

	htmlTemplate := GenerateHtmlTemplate(EmailProps{
		Username: user.Username,
		TaskId:   params.TaskID.Hex(),
		DueDate:  params.DueDate.Local(),
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

	err = models.DeleteReminder(params.Ctx, params.TaskID, params.UserID)
	if err != nil {
		log.Printf("[Reminder] Error Deleting Reminder for %s: %v", params.TaskName, err)
	} else {
		log.Printf("[Reminder] Email sent successfully and reminder deleted for %s", params.TaskName)
	}
}

// GetTasksCron fetches tasks that are due within the next 10 minutes and schedules reminders for them
// This function is intended to be run as a cron job
// It retrieves tasks that are due within the next 10 minutes and schedules reminders for them
func SetTasksCron(ctx context.Context) error {

	collection := db.GetClient().Collection("reminders")

	log.Println(time.Now().Format(time.RFC3339), " - Fetching reminders for cron job")

	var reminders = make([]models.Reminder, 0)

	cursor, err := collection.Find(ctx, bson.M{
		"dueDate": bson.M{
			"$gte": time.Now(),
			"$lte": time.Now().Add(10 * time.Minute),
		},
	})
	if err != nil {
		log.Println("ERROR GETTING REMINDERS CRON: ", err)
		return err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &reminders); err != nil {
		log.Println("ERROR DECODING REMINDERS CRON: ", err)
		return err
	}

	for _, reminder := range reminders {

		payload := ReminderProps{
			TaskID:   reminder.TaskID,
			TaskName: reminder.TaskName,
			UserID:   reminder.UserID,
			DueDate:  reminder.DueDate.Local(),
			Ctx:      ctx,
		}

		err := payload.ScheduleReminderJob()
		if err != nil {
			log.Printf("[Scheduler] Failed to schedule reminder for taskName %v, Error : %v", reminder.TaskName, err)
			continue
		}
	}

	log.Println("✅ total reminders added to cron: ", len(reminders))
	return nil
}
