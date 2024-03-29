package errortype

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	SingleResultErr  = errors.New("single result is nil")
	NotMatchedAnyErr = errors.New("no documents have been matched")
)

type basicQueryInfo struct {
	collection string
	filter     interface{}
	update     interface{}
	doc        interface{}
}

type notFoundError struct {
	basicQueryInfo
}

type duplicatedKeyError struct {
	basicQueryInfo
	error
}

type timeoutError struct {
	basicQueryInfo
	error
}

type internalError struct {
	basicQueryInfo
	error
}

type mongoClientError struct {
	error
}

func (err *basicQueryInfo) setBasicError(col string, filter, update, doc interface{}) {
	err.filter = filter
	err.collection = col
	err.update = update
	err.doc = doc
}

func NotFoundError(col string, filter, update, doc interface{}) error {
	err := &notFoundError{}
	err.setBasicError(col, filter, update, doc)
	return err
}

func DuplicatedKeyError(col string, filter, update, doc interface{}, mongoErr error) error {
	err := &duplicatedKeyError{}
	err.setBasicError(col, filter, update, doc)
	err.error = mongoErr
	return err
}

func TimeoutError(col string, filter, update, doc interface{}, mongoErr error) error {
	err := &timeoutError{}
	err.setBasicError(col, filter, update, doc)
	err.error = mongoErr
	return err
}

func InternalError(col string, filter, update, doc interface{}, mongoErr error) error {
	err := &internalError{}
	err.setBasicError(col, filter, update, doc)
	err.error = mongoErr
	return err
}

func MongoClientError(mongoErr error) error {
	return &mongoClientError{mongoErr}
}

func (e *notFoundError) Error() string {
	return fmt.Sprintf("%s not found. ", e.collection) + getBasicInfoErrorMsg(e.basicQueryInfo)
}

func (e *duplicatedKeyError) Error() string {
	return fmt.Sprintf("%s failed to write due to duplicated key, err: %s ", e.collection, e.error.Error()) + getBasicInfoErrorMsg(e.basicQueryInfo)
}

func (e *timeoutError) Error() string {
	return fmt.Sprintf("%s timeout, err: %s ", e.collection, e.error.Error()) + getBasicInfoErrorMsg(e.basicQueryInfo)
}

func (e *internalError) Error() string {
	return fmt.Sprintf("mongo internal err: %s ", e.error.Error()) + getBasicInfoErrorMsg(e.basicQueryInfo)
}

func (e *mongoClientError) Error() string {
	return fmt.Sprintf("mongo client err: %s ", e.error.Error())
}

func getBasicInfoErrorMsg(e basicQueryInfo) string {
	msg := "| {query info: "
	if e.filter != nil {
		msg += fmt.Sprintf(" filter: %+v", e.filter)
	}

	if e.update != nil {
		msg += fmt.Sprintf(", update: %+v", e.update)
	}

	if e.doc != nil {
		msg += fmt.Sprintf(", doc: %+v", e.doc)
	}
	msg += "}"
	return msg
}
