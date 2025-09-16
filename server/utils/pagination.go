package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GeneratePaginationOptions(page, limit int) *options.FindOptions {
	opts := options.Find()

	opts.SetLimit(int64(limit))
	opts.SetSkip(int64((page - 1) * limit))
	opts.SetSort(bson.M{"created_at": -1})

	return opts
}
