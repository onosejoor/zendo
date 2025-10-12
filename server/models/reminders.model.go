package models

import (
	"context"
	"main/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Reminder struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TaskID     primitive.ObjectID `json:"taskId" bson:"taskId" validate:"required"`
	DueDate    time.Time          `json:"dueDate" bson:"dueDate" validate:"required"`
	UserID     primitive.ObjectID `json:"userId" bson:"userId" validate:"required"`
	TaskName   string             `json:"taskName" bson:"taskName" validate:"required"`
	Expires_At time.Time          `json:"expiresAt" bson:"expiresAt" validate:"required"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
}

func (reminder Reminder) CreateReminder(ctx context.Context) error {
	client := db.GetClient()
	collection := client.Collection("reminders")

	filter := bson.M{
		"taskId": reminder.TaskID,
		"userId": reminder.UserID,
	}

	update := bson.M{
		"$set": bson.M{
			"dueDate":    reminder.DueDate,
			"taskName":   reminder.TaskName,
			"expiresAt":  reminder.Expires_At,
			"created_at": time.Now(),
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func DeleteReminder(ctx context.Context, taskId primitive.ObjectID, userId primitive.ObjectID) error {
	client := db.GetClient()
	collection := client.Collection("reminders")

	_, err := collection.DeleteOne(ctx, bson.M{"taskId": taskId, "userId": userId})

	return err
}

func DeleteAllReminder(ctx context.Context, userId primitive.ObjectID) error {
	client := db.GetClient()
	collection := client.Collection("reminders")

	_, err := collection.DeleteOne(ctx, bson.M{"userId": userId})

	return err
}
