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

type TeamMemberSchema struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TeamID   primitive.ObjectID `json:"team_id" bson:"team_id" validate:"required"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
	Role     string             `json:"role" bson:"role" validate:"required,oneof=owner admin member"`
	JoinedAt time.Time          `json:"joined_at" bson:"joined_at"`
}

type TeamWithRole struct {
	Team     TeamSchema `bson:"team" json:"team"`
	Role     string     `bson:"role" json:"role"`
	JoinedAt time.Time  `bson:"joined_at" json:"joined_at"`
}

type UserWithRole struct {
	User     User      `bson:"user" json:"user"`
	Role     string    `bson:"role" json:"role"`
	JoinedAt time.Time `bson:"joined_at" json:"joined_at"`
}

var teamMembersColl *mongo.Collection

func init() {
	if teamMembersColl == nil {
		teamMembersColl = db.GetClient().Collection("team_members")
	}
}

func (t TeamMemberSchema) CreateTeamMember(ctx context.Context) (*primitive.ObjectID, error) {
	var teamMember TeamMemberSchema

	err := teamMembersColl.FindOne(ctx, bson.M{"user_id": t.UserID, "team_id": t.TeamID}).Decode(&teamMember)
	if err == nil {
		return nil, errors.New("user is already a member of the team")
	}
	if err != mongo.ErrNoDocuments {
		return nil, err
	}

	t.JoinedAt = time.Now()

	id, err := teamMembersColl.InsertOne(ctx, t)
	if err != nil {
		return nil, err
	}
	oid := id.InsertedID.(primitive.ObjectID)
	return &oid, nil
}

func GetTeamMember(ctx context.Context, teamId, userId primitive.ObjectID) (*TeamMemberSchema, error) {
	var member TeamMemberSchema
	err := teamMembersColl.FindOne(ctx, bson.M{
		"team_id": teamId,
		"user_id": userId,
	}).Decode(&member)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func GetTeamsForUser(ctx context.Context, userID primitive.ObjectID, page, limit int) ([]TeamWithRole, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "user_id", Value: userID}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "teams"},
			{Key: "localField", Value: "team_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "team"},
		}}},
		{{Key: "$unwind", Value: "$team"}},
		{{Key: "$project", Value: bson.D{
			{Key: "team", Value: 1},
			{Key: "role", Value: 1},
			{Key: "joined_at", Value: 1},
		}}},
		{{Key: "$skip", Value: (page - 1) * limit}},
		{{Key: "$limit", Value: limit}},
	}

	cursor, err := teamMembersColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results = make([]TeamWithRole, 0)
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetUsersForTeam(ctx context.Context, teamId primitive.ObjectID, page, limit int) ([]UserWithRole, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "team_id", Value: teamId}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "user_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "user"},
		}}},
		{{Key: "$unwind", Value: "$user"}},
		{{Key: "$project", Value: bson.D{
			{Key: "user", Value: 1},
			{Key: "role", Value: 1},
			{Key: "joined_at", Value: 1},
		}}},
		{{Key: "$skip", Value: (page - 1) * limit}},
		{{Key: "$limit", Value: limit}},
	}

	cursor, err := teamMembersColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results = make([]UserWithRole, 0)
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func IsTeamMembers(ctx context.Context, userIds []primitive.ObjectID, teamId primitive.ObjectID, addRole bool) (bool, error) {
	membersColl := db.GetClient().Collection("team_members")

	filter := bson.M{
		"user_id": bson.M{"$in": userIds},
		"team_id": teamId,
	}

	if addRole {
		filter["role"] = bson.M{"$in": []string{"owner", "admin"}}
	}

	number, err := membersColl.CountDocuments(ctx, filter)
	if err != nil {
		log.Println("Error counting IsTeamMember Documents: ", err)
		return false, err
	}

	return number == int64(len(userIds)), nil
}
