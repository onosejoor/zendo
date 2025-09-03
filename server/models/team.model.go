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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TeamSchema struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

var teamColl *mongo.Collection

func init() {
	if teamColl == nil {
		teamColl = db.GetClient().Collection("teams")
	}
}

func GetTeams(ctx context.Context, opts *options.FindOptions) ([]TeamSchema, error) {
	cursor, err := teamColl.Find(ctx, bson.M{}, opts)
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

func GetTeamById(ctx context.Context, teamId, userId primitive.ObjectID) (*TeamSchema, error) {
	count, err := teamMembersColl.CountDocuments(ctx, bson.M{
		"team_id": teamId,
		"user_id": userId,
	})
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, mongo.ErrNoDocuments
	}

	var team TeamSchema
	err = teamColl.FindOne(ctx, bson.M{"_id": teamId}).Decode(&team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (t TeamSchema) CreateTeam(ctx context.Context, userID primitive.ObjectID) (*primitive.ObjectID, error) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	id, err := teamColl.InsertOne(ctx, t)
	if err != nil {
		log.Println("Error Creating Team: ", err.Error())

		return nil, err
	}
	oid := id.InsertedID.(primitive.ObjectID)

	newTeamMember := TeamMemberSchema{
		UserID:   userID,
		TeamID:   oid,
		Role:     "owner",
		JoinedAt: time.Now(),
	}

	_, err = newTeamMember.CreateTeamMember(ctx)
	if err != nil {
		log.Println("Error Adding Team owner to team member schema: ", err.Error())
		return &oid, err
	}

	return &oid, nil
}

func DeleteTeam(ctx context.Context, teamID, userID primitive.ObjectID) error {
	count, err := teamMembersColl.CountDocuments(ctx, bson.M{
		"team_id": teamID,
		"user_id": userID,
		"role":    "owner",
	})
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("only team owner can delete the team")
	}

	_, err = teamColl.DeleteOne(ctx, bson.M{"_id": teamID})
	if err != nil {
		return err
	}

	_, err = teamMembersColl.DeleteMany(ctx, bson.M{"team_id": teamID})
	if err != nil {
		return err
	}

	return nil
}
