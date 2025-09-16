package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateReminderTTLIndex(collection *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expireAt", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Printf("Failed to create TTL index on reminders: %v", err)
	}
}

func CreateModelIndexes(client *mongo.Database) {
	// TeamMembers: index on email, team_id, user_id
	createIndex(client.Collection("team_members"), bson.D{{Key: "email", Value: 1}})
	createIndex(client.Collection("team_members"), bson.D{{Key: "team_id", Value: 1}})
	createIndex(client.Collection("team_members"), bson.D{{Key: "user_id", Value: 1}})

	// Teams: index on name, creator_id
	createIndex(client.Collection("teams"), bson.D{{Key: "name", Value: 1}})
	createIndex(client.Collection("teams"), bson.D{{Key: "creator_id", Value: 1}})

	// Tasks: index on userId, team_id, projectId, status
	createIndex(client.Collection("tasks"), bson.D{{Key: "userId", Value: 1}})
	createIndex(client.Collection("tasks"), bson.D{{Key: "team_id", Value: 1}})
	createIndex(client.Collection("tasks"), bson.D{{Key: "projectId", Value: 1}})
	createIndex(client.Collection("tasks"), bson.D{{Key: "status", Value: 1}})

	// Reminders: index on user_id, team_id, due_date (add TTL if needed)
	createIndex(client.Collection("reminders"), bson.D{{Key: "user_id", Value: 1}})
	createIndex(client.Collection("reminders"), bson.D{{Key: "team_id", Value: 1}})
	CreateReminderTTLIndex(client.Collection("reminders"))

	// TeamInvites: index on email, team_id, status
	createIndex(client.Collection("team_invites"), bson.D{{Key: "email", Value: 1}})
	createIndex(client.Collection("team_invites"), bson.D{{Key: "team_id", Value: 1}})
	createIndex(client.Collection("team_invites"), bson.D{{Key: "status", Value: 1}})
	CreateReminderTTLIndex(client.Collection("team_invites"))

}

func createIndex(collection *mongo.Collection, keys bson.D) {
	indexModel := mongo.IndexModel{Keys: keys}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Printf("Failed to create index %v on %s: %v", keys, collection.Name(), err)
	}
}
