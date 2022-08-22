package rest

import "reflect"

type ApiResponse struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

// Constructor
func NewApiResponse() *ApiResponse {
	res := &ApiResponse{true, SUCCESS.Code, SUCCESS.Message, nil}
	return res
}

func (res *ApiResponse) Succeed(message string, data interface{}) *ApiResponse {
	const isSuccess = true
	r := reflect.ValueOf(&SUCCESS)
	successCode := reflect.Indirect(r).FieldByName("Code")
	resultMessage := reflect.Indirect(r).FieldByName("Message")

	if len(message) == 0 {
		res.Message = resultMessage.String()
	} else {
		res.Message = message
	}

	res.Success = isSuccess
	res.Code = successCode.String()
	res.Data = data

	return res
}

func (res *ApiResponse) Error(codeDesc *CodeDescription, message string, data interface{}) *ApiResponse {
	const isSuccess = false
	r := reflect.ValueOf(&FAILED_INTERNAL_ERROR)
	errorCode := reflect.Indirect(r).FieldByName("Code")
	resultMessage := reflect.Indirect(r).FieldByName("Message")

	if codeDesc != nil {
		res.Code = codeDesc.Code
		res.Message = codeDesc.Message + message
	} else {
		res.Code = errorCode.String()
		res.Message = resultMessage.String()
	}

	res.Success = isSuccess
	res.Data = data

	return res
}
