package utils

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

/**
 * Conversion of interface{} or map[string]interface{} types to protobuf struct
 */
func MapToProtobufStruct(v interface{}) (*structpb.Value, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	s := &structpb.Value{}
	err = protojson.Unmarshal(b, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
