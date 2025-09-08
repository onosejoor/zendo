package models

import (
	"context"
	"main/db"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type TeamInviteSchema struct {
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	TeamID    primitive.ObjectID `json:"team_id" bson:"team_id" validate:"required"`
	ExpiresAt time.Time          `json:"expires_at" bson:"expires_at"`
	Token     string             `json:"token" bson:"token"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

var inviteCollection *mongo.Collection

func init() {
	if inviteCollection == nil {
		inviteCollection = db.GetClient().Collection("team_invites")
	}
}

func (teamMember TeamInviteSchema) CreateMemberInvite(ctx context.Context) error {
	teamMember.CreatedAt = time.Now()
	teamMember.ExpiresAt = time.Now().Add(7 * 24 * time.Hour)

	_, err := inviteCollection.InsertOne(ctx, teamMember)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMemberInvite(ctx context.Context, email string, teamId primitive.ObjectID) error {
	_, err := inviteCollection.DeleteOne(ctx, bson.M{
		"team_id": teamId, "email": email,
	})
	if err != nil {
		return err
	}
	return nil
}

func CheckIfInviteExists(ctx context.Context, email string, teamId primitive.ObjectID) bool {
	emailExists, err := inviteCollection.CountDocuments(ctx, bson.M{
		"email":   email,
		"team_id": teamId,
	})
	if err != nil {
		return false
	}
	if emailExists > 0 {
		return true
	}
	return false
}
