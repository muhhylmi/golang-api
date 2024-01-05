package wrapper

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Data       interface{}
	MetaData   interface{}
	Error      interface{}
	StatusCode string
	Count      int64
}

type BaseWrapperModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Meta    interface{} `json:"meta,omitempty"`
}

func Response(data interface{}, message string, code int, c echo.Context) error {
	success := code < http.StatusBadRequest
	result := BaseWrapperModel{
		Success: success,
		Data:    data,
		Message: message,
		Code:    "0000",
	}

	return c.JSON(code, result)
}

// ResponseError function
func ResponseError(err interface{}, statusCode string, c echo.Context) error {
	errObj := getErrorStatusCode(err)
	result := BaseWrapperModel{
		Success: false,
		Data:    errObj.Data,
		Message: errObj.Message,
		Code:    statusCode,
	}

	return c.JSON(errObj.ResponseCode, result)
}

func getErrorStatusCode(err interface{}) CommonError {
	errData := CommonError{}

	switch obj := err.(type) {
	case RespondError:
		errData.ResponseCode = obj.Code
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	default:
		errData.ResponseCode = http.StatusConflict
		errData.Code = http.StatusConflict
		return errData
	}
}

func ResultSuccess(data interface{}) Result {
	return Result{
		Data: data,
	}
}

func ResultFailed(err interface{}, statusCode string) Result {
	return Result{
		Data:       nil,
		Error:      err,
		StatusCode: statusCode,
	}
}
