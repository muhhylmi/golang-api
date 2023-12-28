package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Data     interface{}
	MetaData interface{}
	Error    interface{}
	Count    int64
}

type BaseWrapperModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Meta    interface{} `json:"meta,omitempty"`
}

func Response(data interface{}, message string, code int, c echo.Context) error {
	success := false

	if code < http.StatusBadRequest {
		success = true
	}

	result := BaseWrapperModel{
		Success: success,
		Data:    data,
		Message: message,
		Code:    code,
	}

	return c.JSON(code, result)
}

// ResponseError function
func ResponseError(err interface{}, c echo.Context) error {
	errObj := getErrorStatusCode(err)
	result := BaseWrapperModel{
		Success: false,
		Data:    errObj.Data,
		Message: errObj.Message,
		Code:    errObj.Code,
	}

	return c.JSON(errObj.ResponseCode, result)
}

func getErrorStatusCode(err interface{}) CommonError {
	errData := CommonError{}

	switch obj := err.(type) {
	case RespondError:
		errData.ResponseCode = http.StatusBadRequest
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
