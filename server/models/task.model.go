package models

import (
	"context"
	"main/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	ID          primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string              `json:"title" bson:"title" validate:"required"`
	Description string              `json:"description" bson:"description" validate:"required"`
	UserId      primitive.ObjectID  `json:"userId" bson:"userId"`
	SubTasks    []SubTask           `json:"subTasks,omitempty" bson:"subTasks,omitempty"`
	ProjectId   *primitive.ObjectID `json:"projectId,omitempty" bson:"projectId,omitempty"`
	DueDate     time.Time           `json:"dueDate" bson:"dueDate" validate:"required"`
	Status      string              `json:"status" bson:"status" validate:"required"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
}

type SubTask struct {
	ID        string `json:"_id" bson:"_id"`
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}

func CreateTask(p Task, ctx context.Context, userId primitive.ObjectID) (id any, err error) {

	if p.ProjectId != nil {
		id, err := p.CreateTaskWithTransaction(ctx, userId)
		if err != nil {
			return nil, err
		}
		return id, nil
	}

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
	now := time.Now().Local().Truncate(24 * time.Hour)

	for _, task := range t {
		dueDate := task.DueDate.Local().Truncate(24 * time.Hour)

		if dueDate.Equal(now) && task.Status != "completed" {
			isDueToday++
		}

		if task.Status == "completed" {
			completedTasks++
		}
	}

	var completionRate int
	if len(t) > 0 {
		completionRate = int(float64(completedTasks) / float64(len(t)) * 100)
	}

	return isDueToday, completionRate, completedTasks
}

func (task Task) CreateTaskWithTransaction(ctx context.Context, userId primitive.ObjectID) (primitive.ObjectID, error) {
	client := db.GetClientWithoutDB()

	session, err := client.StartSession()
	if err != nil {
		return primitive.NilObjectID, err
	}
	defer session.EndSession(ctx)

	tasksCollection := client.Database("zendo").Collection("tasks")
	projectsCollection := client.Database("zendo").Collection("projects")

	callback := func(sessCtx mongo.SessionContext) (any, error) {
		id, err := tasksCollection.InsertOne(sessCtx, Task{
			Title:       task.Title,
			Description: task.Description,
			UserId:      userId,
			SubTasks:    task.SubTasks,
			ProjectId:   task.ProjectId,
			DueDate:     task.DueDate,
			Status:      task.Status,
			CreatedAt:   time.Now(),
		})
		if err != nil {
			return nil, err
		}

		// Step 2: Update the project's total task count
		update := bson.M{"$inc": bson.M{"totalTasks": 1}}
		filter := bson.M{"_id": task.ProjectId}

		_, err = projectsCollection.UpdateOne(sessCtx, filter, update)
		if err != nil {
			return nil, err
		}

		return id.InsertedID, nil
	}

	id, err := session.WithTransaction(ctx, callback)
	return id.(primitive.ObjectID), err
}

func (task Task) DeleteTaskWithTransaction(ctx context.Context) error {
	client := db.GetClientWithoutDB()
	id := task.ID

	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	tasksCollection := client.Database("zendo").Collection("tasks")
	projectsCollection := client.Database("zendo").Collection("projects")

	callback := func(sessCtx mongo.SessionContext) (any, error) {
		_, err := tasksCollection.DeleteOne(sessCtx, bson.M{"_id": id})
		if err != nil {
			return nil, err
		}

		update := bson.M{"$inc": bson.M{"totalTasks": -1}}
		filter := bson.M{"_id": task.ProjectId}

		_, err = projectsCollection.UpdateOne(sessCtx, filter, update)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	return err
}

func DeleteAllTasksWithTransaction(ctx context.Context, user *UserRes) error {
	client := db.GetClientWithoutDB()

	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	tasksCollection := client.Database("zendo").Collection("tasks")
	projectsCollection := client.Database("zendo").Collection("projects")

	callback := func(sessCtx mongo.SessionContext) (any, error) {
		cursor, err := tasksCollection.Find(sessCtx, bson.M{"userId": user.ID})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(sessCtx)

		projectTaskCount := make(map[primitive.ObjectID]int)
		for cursor.Next(sessCtx) {
			var task Task
			if err := cursor.Decode(&task); err != nil {
				return nil, err
			}
			if task.ProjectId != nil {
				projectTaskCount[*task.ProjectId]++
			}
		}

		_, err = tasksCollection.DeleteMany(sessCtx, bson.M{"userId": user.ID})
		if err != nil {
			return nil, err
		}

		for projectID, count := range projectTaskCount {
			filter := bson.M{"_id": projectID}
			update := bson.M{"$inc": bson.M{"totalTasks": -count}}

			_, err := projectsCollection.UpdateOne(sessCtx, filter, update)
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	return err
}
