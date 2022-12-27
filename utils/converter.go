package utils

import (
	"encoding/json"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
 * Conversion of string types to objectID
 */
func MapToObjectID(v string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(v)
	if err != nil {
		log.Println("Invalid ID")
	}
	return objectID, err
}

/**
 * Conversion of objectID to string types
 */
func MapToStringID(v primitive.ObjectID) string {
	return v.Hex()
}

/**
 * Conversion of bson.D to struct
 */
func MapToStruct(v interface{}, target interface{}) *interface{} {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &target)
	if err != nil {
		panic(err)
	}

	return &target
}

/**
 * Conversion of supplied struct to a pointer via interface{}
 */
func MapToSructPointer(obj interface{}) interface{} {
	vp := reflect.New(reflect.TypeOf(obj))
	vp.Elem().Set(reflect.ValueOf(obj))
	return vp.Interface()
}
