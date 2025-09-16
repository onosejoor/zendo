package models

import (
	"context"
	"errors"
	"main/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamMemberSchema struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email" validate:"required"`
	TeamID   primitive.ObjectID `json:"team_id" bson:"team_id" validate:"required"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
	Role     string             `json:"role" bson:"role" validate:"required,oneof=owner admin member"`
	JoinedAt time.Time          `json:"joined_at" bson:"joined_at"`
}

type TeamWithRole struct {
	TeamSchema  `bson:"team" json:",inline"`
	Role        string    `bson:"role" json:"role"`
	MemberCount int       `bson:"members_count" json:"members_count"`
	JoinedAt    time.Time `bson:"joined_at" json:"joined_at"`
}

type UserWithRole struct {
	User     `bson:"user" json:",inline"`
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
	err := teamMembersColl.FindOne(ctx, bson.M{"user_id": t.UserID, "team_id": t.TeamID}).Err()
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

func GetTeamMembersRaw(ctx context.Context, teamId primitive.ObjectID) (*[]TeamMemberSchema, error) {
	var members []TeamMemberSchema
	cursor, err := teamMembersColl.Find(ctx, bson.M{
		"team_id": teamId,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &members); err != nil {
		return nil, err
	}

	return &members, nil
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

func GetTeamMemberByEmail(ctx context.Context, teamId primitive.ObjectID, userEmail string) (*TeamMemberSchema, error) {
	var member TeamMemberSchema
	err := teamMembersColl.FindOne(ctx, bson.M{
		"team_id": teamId,
		"email":   userEmail,
	}).Decode(&member)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func DeleteTeamMember(memberId, teamId primitive.ObjectID, ctx context.Context) error {
	_, err := teamMembersColl.DeleteOne(ctx, bson.M{
		"team_id": teamId,
		"user_id": memberId,
	})
	if err != nil {
		return err
	}

	_, err = tasksCollection.UpdateMany(
		ctx,
		bson.M{
			"team_id":   teamId,
			"assignees": bson.M{"$in": []primitive.ObjectID{memberId}},
		},
		bson.M{
			"$pull": bson.M{"assignees": memberId},
		},
	)
	return err
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
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "team_members"},
			{Key: "localField", Value: "team_id"},
			{Key: "foreignField", Value: "team_id"},
			{Key: "as", Value: "members"},
		}}},
		{{Key: "$addFields", Value: bson.D{
			{Key: "members_count", Value: bson.D{{Key: "$size", Value: "$members"}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "team", Value: 1},
			{Key: "role", Value: 1},
			{Key: "joined_at", Value: 1},
			{Key: "members_count", Value: 1},
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

func IsTeamMembers(ctx context.Context, userIds []primitive.ObjectID, teamId primitive.ObjectID, restrictToAdmins bool) (bool, string, error) {
	membersColl := db.GetClient().Collection("team_members")

	filter := bson.M{
		"user_id": bson.M{"$in": userIds},
		"team_id": teamId,
	}

	if restrictToAdmins {
		filter["role"] = bson.M{"$in": []string{"owner", "admin"}}
	}

	var member TeamMemberSchema
	err := membersColl.FindOne(ctx, filter).Decode(&member)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, "", nil
		}
		return false, "", err
	}

	return true, member.Role, nil
}
