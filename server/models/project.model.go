package models

import (
	"context"
	"log"
	"main/db"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
