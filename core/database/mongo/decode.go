package database_mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func DecodeCursor[T any](cursor *mongo.Cursor) ([]T, error) {
	defer cursor.Close(context.Background())
	var slice []T
	for cursor.Next(context.Background()) {
		var doc T
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		slice = append(slice, doc)
	}
	return slice, nil
}
