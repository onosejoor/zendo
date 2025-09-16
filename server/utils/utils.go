package utils

import (
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HexToObjectID(hex string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return primitive.NilObjectID
	}
	return oid
}

var isPulledEnvs = false

func PullEnv() {
	if isPulledEnvs {
		return
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("No .env file found, using environment variables")
	}
	isPulledEnvs = true

}
