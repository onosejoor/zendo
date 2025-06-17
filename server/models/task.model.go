package models

import (
	"context"
	"main/db"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	SubTasks    []SubTasks         `json:"subTasks" bson:"subTasks" validate:"required"`
	DueDate     time.Time          `json:"dueDate" bson:"dueDate" validate:"required"`
	Status      string             `json:"status" bson:"status" validate:"required"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

type SubTasks struct {
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}

func CreateTask(p Task, ctx context.Context, userId primitive.ObjectID) (id any, err error) {
	client := db.GetClient()
	collection := client.Collection("tasks")

	newTaskId, err := collection.InsertOne(ctx, Task{
		Title:       p.Title,
		Description: p.Description,
		UserId:      userId,
		SubTasks:    p.SubTasks,
		DueDate:     p.DueDate,
		Status:      p.Status,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return newTaskId.InsertedID, nil
}
