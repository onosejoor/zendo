package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTeamMemberEmailIndex(collection *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}},
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Failed to create email index on team_members: %v", err)
	}
}

func CreateTeamNameIndex(collection *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "name", Value: 1}},
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Failed to create name index on teams: %v", err)
	}
}

func CreateTaskUserIdIndex(collection *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "userId", Value: 1}},
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Failed to create userId index on tasks: %v", err)
	}
}

func CreateReminderTTLIndex(collection *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiresAt", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Failed to create TTL index: %v", err)
	}
}

func CreateDbIndex(client *mongo.Database) {
	CreateReminderTTLIndex(client.Collection("reminders"))
	CreateReminderTTLIndex(client.Collection("team_invites"))
	CreateTeamMemberEmailIndex(client.Collection("team_members"))
	CreateTeamNameIndex(client.Collection("teams"))
	CreateTaskUserIdIndex(client.Collection("tasks"))
}
