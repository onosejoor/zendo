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
	ID          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string               `json:"title" bson:"title" validate:"required"`
	Description string               `json:"description" bson:"description"`
	UserId      primitive.ObjectID   `json:"userId" bson:"userId"`
	SubTasks    []SubTask            `json:"subTasks,omitempty" bson:"subTasks,omitempty"`
	ProjectId   primitive.ObjectID   `json:"projectId,omitempty" bson:"projectId,omitempty"`
	TeamID      primitive.ObjectID   `json:"team_id,omitempty" bson:"team_id,omitempty"`
	DueDate     time.Time            `json:"dueDate" bson:"dueDate" validate:"required"`
	Status      string               `json:"status" bson:"status" validate:"required"`
	Assignees   []primitive.ObjectID `json:"assignees" bson:"assignees"`
	CreatedAt   time.Time            `json:"created_at" bson:"created_at"`
}

type AssigneeInfo struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Avatar   string             `bson:"avatar" json:"avatar"`
}

type TaskWithAssignees struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	TeamID      primitive.ObjectID `bson:"team_id" json:"team_id"`
	DueDate     time.Time          `bson:"dueDate" json:"dueDate"`
	Status      string             `bson:"status" json:"status"`
	SubTask     []SubTask          `bson:"subTasks" json:"subTasks"`
	Assignees   []AssigneeInfo     `bson:"assignee_details" json:"assignees"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
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

	p.UserId = userId
	p.CreatedAt = time.Now()

	newTaskId, err := tasksCollection.InsertOne(ctx, p)
	if err != nil {
		return nil, err
	}

	return newTaskId.InsertedID, nil
}

func CheckAssignee(assignees []primitive.ObjectID, teamId primitive.ObjectID, ctx context.Context) (bool, error) {
	isMembers, _, err := IsTeamMembers(ctx, assignees, teamId, false)
	if err != nil {
		return false, err
	}
	if isMembers {
		return true, nil
	}

	return false, nil

}

func GetTaskReminderSent(taskId primitive.ObjectID, client *mongo.Database, ctx context.Context) (bool, error) {
	remindersCollection := client.Collection("reminders")

	completedCount, err := tasksCollection.CountDocuments(
		ctx,
		bson.M{"_id": taskId, "status": "completed"},
	)
	if err != nil {
		return true, err
	}

	if completedCount > 0 {
		return true, nil
	}

	reminderCount, err := remindersCollection.CountDocuments(
		ctx,
		bson.M{"taskId": taskId},
	)
	if err != nil {
		return true, err
	}

	if reminderCount > 0 {
		return false, nil
	}

	return true, nil
}

func GetTasksForTeam(ctx context.Context, teamId primitive.ObjectID, page, limit int) ([]TaskWithAssignees, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "team_id", Value: teamId}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "assignees"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "assignee_details"},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "title", Value: 1},
			{Key: "description", Value: 1},
			{Key: "team_id", Value: 1},
			{Key: "dueDate", Value: 1},
			{Key: "status", Value: 1},
			{Key: "assignees", Value: 1},
			{Key: "assignee_details", Value: bson.D{
				{Key: "username", Value: 1},
				{Key: "avatar", Value: 1},
				{Key: "email", Value: 1},
				{Key: "_id", Value: 1},
			}},
		}}},
		{{Key: "$skip", Value: (page - 1) * limit}},
		{{Key: "$limit", Value: limit}},
	}
	cursor, err := tasksCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	var tasks []TaskWithAssignees
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskForTeamById(ctx context.Context, teamId, taskId primitive.ObjectID) (*TaskWithAssignees, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "team_id", Value: teamId},
			{Key: "_id", Value: taskId},
		}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "assignees"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "assignee_details"},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "title", Value: 1},
			{Key: "description", Value: 1},
			{Key: "team_id", Value: 1},
			{Key: "dueDate", Value: 1},
			{Key: "subTasks", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "status", Value: 1},
			{Key: "assignees", Value: 1},
			{Key: "assignee_details", Value: bson.D{
				{Key: "username", Value: 1},
				{Key: "email", Value: 1},
				{Key: "_id", Value: 1},
				{Key: "avatar", Value: 1},
			}},
		}}},
	}

	cursor, err := tasksCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var task TaskWithAssignees
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		return &task, nil
	}

	// No task found
	return nil, mongo.ErrNoDocuments
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
		task.UserId = userId
		task.CreatedAt = time.Now()
		id, err := tasksCollection.InsertOne(sessCtx, task)
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
