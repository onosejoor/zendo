package models

import (
	"context"
	"main/db"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string              `json:"title" bson:"title" validate:"required"`
	Description string              `json:"description" bson:"description" validate:"required"`
	UserId      primitive.ObjectID  `json:"userId" bson:"userId"`
	SubTasks    []SubTasks          `json:"subTasks,omitempty" bson:"subTasks,omitempty"`
	ProjectId   *primitive.ObjectID `json:"projectId,omitempty" bson:"projectId,omitempty"`
	DueDate     time.Time           `json:"dueDate" bson:"dueDate" validate:"required"`
	Status      string              `json:"status" bson:"status" validate:"required"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
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
		ProjectId:   p.ProjectId,
		DueDate:     p.DueDate,
		Status:      p.Status,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return newTaskId.InsertedID, nil
}

func GetCompletionRateAndDueDate(t []Task) (int, int, int) {
	isDueToday := 0
	completedTasks := 0

	for _, task := range t {
		if time.Time.Equal(task.DueDate, time.Now()) {
			isDueToday++
		}

		if task.Status == "completed" {
			completedTasks++
		}
	}

	completionRate := (completedTasks / len(t)) * 100

	return isDueToday, completionRate, completedTasks

}
