package models

import (
	"context"
	"main/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TeamSchema struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

var client *mongo.Collection

func init() {
	if client == nil {
		client = db.GetClient().Collection("teams")
	}
}

func GetTeams(ctx context.Context, opts *options.FindOptions) ([]TeamSchema, error) {
	cursor, err := client.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var teams []TeamSchema
	if err := cursor.All(ctx, &teams); err != nil {
		return nil, err
	}
	return teams, nil

}

func (t TeamSchema) CreateTeam(ctx context.Context) (*primitive.ObjectID, error) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	id, err := client.InsertOne(ctx, t)
	if err != nil {
		return nil, err
	}
	oid := id.InsertedID.(primitive.ObjectID)
	return &oid, nil
}
