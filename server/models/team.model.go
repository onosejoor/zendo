package models

import (
	"context"
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
	CreatorId   primitive.ObjectID `json:"creator_id" bson:"creator_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type TeamWithMemberAndRole struct {
	TeamSchema   `json:",inline" bson:",inline"`
	MembersCount int       `json:"members_count" bson:"members_count"`
	Role         string    `json:"role" bson:"role"`
	JoinedAt     time.Time `json:"joined_at" bson:"joined_at"`
}

var teamColl *mongo.Collection

func init() {
	if teamColl == nil {
		teamColl = db.GetClient().Collection("teams")
	}
}

func CheckTeamExist(teamId primitive.ObjectID, ctx context.Context) bool {

	number, err := teamColl.CountDocuments(ctx, bson.M{
		"_id": teamId,
	})
	if err != nil {
		log.Println("ERROR CHECKING TEAM EXIST: ", err)
		return false
	}

	return number > 0

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

func GetTeamById(ctx context.Context, teamId, userId primitive.ObjectID) (*TeamWithMemberAndRole, error) {
	var member TeamMemberSchema
	err := teamMembersColl.FindOne(ctx, bson.M{
		"team_id": teamId,
		"user_id": userId,
	}).Decode(&member)
	if err != nil {
		return nil, err
	}

	numberOfMembers, err := teamMembersColl.CountDocuments(ctx, bson.M{
		"team_id": teamId,
	})
	if err != nil {
		return nil, err
	}

	var team TeamWithMemberAndRole
	err = teamColl.FindOne(ctx, bson.M{"_id": teamId}).Decode(&team)
	if err != nil {
		return nil, err
	}

	team.MembersCount = int(numberOfMembers)
	team.Role = member.Role
	team.JoinedAt = member.JoinedAt

	return &team, nil
}

func (t TeamSchema) CreateTeam(ctx context.Context, userID primitive.ObjectID) (*primitive.ObjectID, error) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	t.CreatorId = userID

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

	_, err := teamColl.DeleteOne(ctx, bson.M{"_id": teamID})
	if err != nil {
		return err
	}

	_, err = teamMembersColl.DeleteMany(ctx, bson.M{"team_id": teamID})
	if err != nil {
		return err
	}

	_, err = tasksCollection.DeleteMany(ctx, bson.M{"team_id": teamID})
	if err != nil {
		return err
	}

	return nil
}

func CheckMemberRoleMatch(userId, teamId primitive.ObjectID, ctx context.Context, roles []string) bool {
	memberExist := teamMembersColl.FindOne(ctx, bson.M{
		"user_id": userId, "team_id": teamId, "role": bson.M{"$in": roles},
	}).Err() == nil

	return memberExist
}
