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

type TeamInviteSchema struct {
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	TeamID    primitive.ObjectID `json:"team_id" bson:"team_id" validate:"required"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	ExpiresAt time.Time          `json:"expiresAt" bson:"expiresAt"`
	Token     string             `json:"token" bson:"token"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

var inviteCollection *mongo.Collection

func init() {
	if inviteCollection == nil {
		inviteCollection = db.GetClient().Collection("team_invites")
	}
}
func (invite *TeamInviteSchema) CreateOrUpdateInvite(ctx context.Context) error {
	filter := bson.M{
		"email":   invite.Email,
		"team_id": invite.TeamID,
	}

	update := bson.M{
		"$set": bson.M{
			"status":    invite.Status,
			"token":     invite.Token,
			"updatedAt": time.Now(),
			"expiresAt": time.Now().Add(7 * 24 * time.Hour),
		},
		"$setOnInsert": bson.M{
			"user_id":   invite.UserID,
			"createdAt": time.Now(),
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := inviteCollection.UpdateOne(ctx, filter, update, opts)
	return err
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

func CheckIfInviteExists(ctx context.Context, email string, teamId primitive.ObjectID) *TeamInviteSchema {
	var invite TeamInviteSchema
	err := inviteCollection.FindOne(ctx, bson.M{
		"email":   email,
		"team_id": teamId,
	}).Decode(&invite)

	if err != nil {
		return nil
	}

	return &invite
}

func DeleteInvite(email string, teamId primitive.ObjectID, ctx context.Context) error {
	_, err := inviteCollection.DeleteOne(ctx, bson.M{"email": email, "team_id": teamId})
	return err
}
