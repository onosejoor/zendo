package models

import (
	"context"
	"errors"
	"log"
	"main/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Project struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	OwnerID     primitive.ObjectID `json:"ownerId" bson:"ownerId"`
	TotalTasks  int                `json:"totalTasks" bson:"totalTasks"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

func CreateProject(p Project, userId primitive.ObjectID, ctx context.Context) (id primitive.ObjectID, err error) {
	collection := db.GetClient().Collection("projects")

	newProject, err := collection.InsertOne(ctx, Project{
		Name:        p.Name,
		Description: p.Description,
		OwnerID:     userId,
		TotalTasks:  0,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		log.Println("Error inserting project: ", err)
		return primitive.NilObjectID, err
	}

	return newProject.InsertedID.(primitive.ObjectID), nil

}

type DeleteData struct {
	TotalTasksDeleted    int64 `json:"total_tasks_deleted"`
	TotalProjectsDeleted int64 `json:"total_projects_deleted"`
}

func DeleteAllProjectsWithTransaction(ctx context.Context, user *UserRes) (DeleteData, error) {
	client := db.GetClientWithoutDB()

	session, err := client.StartSession()
	if err != nil {
		return DeleteData{}, err
	}
	defer session.EndSession(ctx)

	tasksCollection := client.Database("zendo").Collection("tasks")
	projectsCollection := client.Database("zendo").Collection("projects")

	callback := func(sessCtx mongo.SessionContext) (any, error) {
		cursor, err := projectsCollection.Find(sessCtx, bson.M{"ownerId": user.ID})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(sessCtx)

		var projects []Project
		if err := cursor.All(sessCtx, &projects); err != nil {
			log.Println("Error parsing db data to slice:", err)
			return nil, err
		}

		if len(projects) == 0 {
			return DeleteData{
				TotalTasksDeleted:    0,
				TotalProjectsDeleted: 0,
			}, nil
		}

		projectIDs := make([]primitive.ObjectID, len(projects))
		for i, p := range projects {
			projectIDs[i] = p.ID
		}

		tasksDeleteResult, err := tasksCollection.DeleteMany(sessCtx, bson.M{
			"projectId": bson.M{"$in": projectIDs},
		})
		if err != nil {
			return nil, err
		}

		projectsDeleteResult, err := projectsCollection.DeleteMany(sessCtx, bson.M{
			"ownerId": user.ID,
		})
		if err != nil {
			return nil, err
		}

		return DeleteData{
			TotalTasksDeleted:    tasksDeleteResult.DeletedCount,
			TotalProjectsDeleted: projectsDeleteResult.DeletedCount,
		}, nil
	}

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		return DeleteData{}, err
	}

	data, ok := result.(DeleteData)
	if !ok {
		return DeleteData{}, errors.New("unexpected result type from transaction")
	}

	return data, nil
}
