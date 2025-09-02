package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func HexToObjectID(hex string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return primitive.NilObjectID
	}
	return oid
}
