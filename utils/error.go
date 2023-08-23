package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CommonError struct
type CommonError struct {
	Code         int         `json:"code"`
	ResponseCode int         `json:"responseCode,omitempty"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

// BadRequest struct
type BadRequest struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewBadRequest
func NewBadRequest() BadRequest {
	errObj := BadRequest{}
	errObj.Message = "Bad Request"
	errObj.Code = http.StatusBadRequest

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

func BindValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := c.Validate(i); err != nil {
		return err
	}
	return nil
}
