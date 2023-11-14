package utils

import (
	"net/http"
)

// CommonError struct
type CommonError struct {
	Code         int         `json:"code"`
	ResponseCode int         `json:"responseCode,omitempty"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

// BadRequest struct
type RespondError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewBadRequest
func NewBadRequest(message string) RespondError {
	errObj := RespondError{}
	errObj.Message = message
	errObj.Code = http.StatusBadRequest

	return errObj
}

func NewNotFound(message string) RespondError {
	errObj := RespondError{}
	errObj.Message = message
	errObj.Code = http.StatusNotFound

	return errObj
}

// Conflict struct
type Conflict struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewConflict() Conflict {
	errObj := Conflict{}
	errObj.Message = "Conflict"
	errObj.Code = http.StatusConflict

	return errObj
}
