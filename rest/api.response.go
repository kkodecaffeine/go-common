package rest

import (
	"reflect"

	"github.com/kkodecaffeine/go-common/errorcode"
)

type ApiResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

// Constructor
func NewApiResponse() *ApiResponse {
	res := &ApiResponse{errorcode.SUCCESS.Code, errorcode.SUCCESS.Message, nil}

	return res
}

func (res *ApiResponse) Succeed(message string, data interface{}) *ApiResponse {
	r := reflect.ValueOf(&errorcode.SUCCESS)
	successCode := reflect.Indirect(r).FieldByName("Code")
	resultMessage := reflect.Indirect(r).FieldByName("Message")

	if len(message) == 0 {
		res.Message = resultMessage.String()
	} else {
		res.Message = message
	}

	res.Code = successCode.String()
	res.Data = data

	return res
}

func (res *ApiResponse) Created(message string, data interface{}) *ApiResponse {
	r := reflect.ValueOf(&errorcode.CREATED)
	successCode := reflect.Indirect(r).FieldByName("Code")
	resultMessage := reflect.Indirect(r).FieldByName("Message")

	if len(message) == 0 {
		res.Message = resultMessage.String()
	} else {
		res.Message = message
	}

	res.Code = successCode.String()
	res.Data = data

	return res
}

func (res *ApiResponse) Error(codeDesc *errorcode.CodeDescription, message string, data interface{}) *ApiResponse {
	r := reflect.ValueOf(&errorcode.FAILED_INTERNAL_ERROR)
	errorCode := reflect.Indirect(r).FieldByName("Code")
	resultMessage := reflect.Indirect(r).FieldByName("Message")

	if codeDesc != nil {
		res.Code = codeDesc.Code
		res.Message = codeDesc.Message + message
	} else {
		res.Code = errorCode.String()
		res.Message = resultMessage.String()
	}

	res.Data = data

	return res
}

type CustomError struct {
	HttpStatusCode int
	CodeDesc       *errorcode.CodeDescription
	Message        string
	Data           interface{}
}

func (e *CustomError) Error() *CustomError {
	var result *CustomError

	result.HttpStatusCode = e.HttpStatusCode
	result.CodeDesc = e.CodeDesc
	result.CodeDesc.Message = e.CodeDesc.Message + e.Message

	return result
}
