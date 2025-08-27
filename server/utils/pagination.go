package utils

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GeneratePaginationOptions(page, limit string) *options.FindOptions {
	if page == "" {
		page = "1"
	}
	if limit == "" {
		page = "10"
	}

	nextPage, pageErr := strconv.Atoi(page)
	pageLimit, limitErr := strconv.Atoi(limit)
	if pageErr != nil || limitErr != nil {
		return nil
	}
	opts := options.Find()

	opts.SetLimit(int64(pageLimit))
	opts.SetSkip(int64((nextPage - 1) * pageLimit))
	opts.SetSort(bson.M{"created_at": -1})

	return opts
}
