package models

import (
	"context"
	"main/db"
	"main/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" validate:"required"`
	Description string             `json:"description" bson:"description"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	SubTasks    []SubTask          `json:"subTasks,omitempty" bson:"subTasks,omitempty"`
	ProjectId   primitive.ObjectID `json:"projectId,omitempty" bson:"projectId,omitempty"`
	TeamID      primitive.ObjectID `json:"team_id,omitempty" bson:"team_id,omitempty"`
	DueDate     time.Time          `json:"dueDate" bson:"dueDate" validate:"required"`
	Status      string             `json:"status" bson:"status" validate:"required"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

var (
	tasksCollection    *mongo.Collection
	projectsCollection *mongo.Collection
)

func init() {
	var cl = db.GetClient()
	if tasksCollection == nil {
		tasksCollection = cl.Collection("tasks")
	}
	if projectsCollection == nil {
		projectsCollection = cl.Collection("projects")
	}
}

type SubTask struct {
	ID        string `json:"_id" bson:"_id"`
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}

func CreateTask(p Task, ctx context.Context, userId primitive.ObjectID) (id any, err error) {
	if p.ProjectId != primitive.NilObjectID {
		id, err := p.CreateTaskWithTransaction(ctx, userId)
		if err != nil {
			return nil, err
		}
		return id, nil
	}

	newTaskId, err := tasksCollection.InsertOne(ctx, Task{
		Title:       p.Title,
		Description: p.Description,
		UserId:      userId,
		SubTasks:    p.SubTasks,
		DueDate:     p.DueDate,
		TeamID:      p.TeamID,
		Status:      p.Status,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return newTaskId.InsertedID, nil
}

func GetTaskReminderSent(taskId primitive.ObjectID, client *mongo.Database, ctx context.Context) (bool, error) {
	remindersCollection := client.Collection("reminders")

	var taskResult bson.M
	taskProjection := bson.M{"_id": 0, "status": 1}

	err := tasksCollection.FindOne(
		ctx,
		bson.M{"_id": taskId}, options.FindOne().SetProjection(taskProjection),
	).Decode(&taskResult)
	if err != nil {
		return true, err
	}

	if taskResult["status"] == "completed" {
		return true, nil
	}

	var result bson.M

	reminderProjection := bson.M{"_id": 0}

	opts := options.FindOne().SetProjection(reminderProjection)

	err = remindersCollection.FindOne(
		ctx,
		bson.M{"taskId": taskId}, opts,
	).Decode(&result)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return true, err
		}
		return false, err
	}
	return false, nil
}

func GetTasksForTeam(ctx context.Context, teamId primitive.ObjectID, page, limit int) ([]Task, error) {
	opts := utils.GeneratePaginationOptions(page, limit)
	cursor, err := tasksCollection.Find(ctx, bson.M{"team_id": teamId}, opts)
	if err != nil {
		return nil, err
	}

	var tasksSlice []Task
	if err := cursor.All(ctx, &tasksSlice); err != nil {
		return nil, err
	}
	return tasksSlice, nil
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
			if task.ProjectId != primitive.NilObjectID {
				projectTaskCount[task.ProjectId]++
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
