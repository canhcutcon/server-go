package hepper

import "server-go/internal/pkg/validator"

type HttpResponse struct {
	Result        interface{}                 `json:"result"`
	Success       bool                        `json:"success"`
	StatusCode    int                         `json:"status_code"`
	ValidateError *[]validator.ValidatorError `json:"validate_error"`
	Error         interface{}                 `json:"error"`
}

func HttpResponseError(result interface{}, success bool, statusCode int, validateError *[]validator.ValidatorError, err interface{}) *HttpResponse {
	return &HttpResponse{
		Result:        result,
		Success:       success,
		StatusCode:    statusCode,
		ValidateError: validateError,
		Error:         err,
	}
}

func HttpResponseSuccess(result interface{}, success bool, statusCode int) *HttpResponse {
	return &HttpResponse{
		Result:        result,
		Success:       success,
		StatusCode:    statusCode,
		ValidateError: nil,
		Error:         nil,
	}
}
