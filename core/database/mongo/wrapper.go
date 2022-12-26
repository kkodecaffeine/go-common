package database_mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection[T any] struct {
	*mongo.Collection
}

func NewCollection[T any](client *mongo.Client, databaseName, collectionName string) *Collection[T] {
	coll := client.Database(databaseName).Collection(collectionName)
	return &Collection[T]{Collection: coll}
}

// InsertOne adds a document to the database into collection
func (coll *Collection[T]) InsertOne(doc bson.D) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := coll.Collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, fmt.Errorf("could not create a document into : %s with error: %q", coll.Name(), err)
	}
	id := res.InsertedID
	return id, nil
}

// FindAll returns items from the database
func (coll *Collection[T]) FindAll(filter interface{}, opts ...*options.FindOptions) ([]T, error) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	cursor, err := coll.Find(ctx, filter, opts...)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) || errors.Is(err, errors.New("no documents have been matched")) {
			return nil, fmt.Errorf("could not find a document : %q with error: %q", filter, err)
		}

		if mongo.IsDuplicateKeyError(err) {
			return nil, fmt.Errorf("duplicate key error collection : %q with error: %q", filter, err)
		}

		if mongo.IsTimeout(err) || errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("timeout : %q with error: %q", filter, err)
		}

		return nil, fmt.Errorf("internal error : %q with error: %q", filter, err)
	}

	resultSlice, err := DecodeCursor[T](cursor)
	if err != nil {
		return nil, fmt.Errorf("decode error : %q with error: %q", filter, err)
	}
	return resultSlice, nil
}

// FindOne returns a single item from the database
func (coll *Collection[T]) FindOne(data, filter interface{}) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	singleResult := coll.Collection.FindOne(ctx, filter)
	if err := decode(singleResult, data); err != nil {
		if err == mongo.ErrNoDocuments || errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("could not find a document : %q with error: %q", filter, err)
		}
		return nil, fmt.Errorf("an error:%q occured while finding filter : %q", err, filter)
	}

	return singleResult, nil
}

func decode(result *mongo.SingleResult, v interface{}) error {
	if result == nil {
		return nil
	}
	if err := result.Decode(v); err != nil {
		return err
	}
	return nil
}
