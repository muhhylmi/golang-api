package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mockUsecase "golang-api/mocks/modules/books/usecases"
	"golang-api/utils/logger"
	"golang-api/utils/wrapper"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllBook(t *testing.T) {
	type TestCase struct {
		name           string
		expectHttpCode int
		response       error
	}

	// create handler
	testCases := []TestCase{
		{
			name:           "should return error",
			expectHttpCode: http.StatusConflict,
		},
		{
			name:           "should return success",
			expectHttpCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create mocks
			l := logger.Newlogger()
			uc := mockUsecase.NewUsecases(t)
			h := HTTPHandler{
				Logger: l,
			}
			app := echo.New()
			req := httptest.NewRequest("GET", "/books", nil)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := app.NewContext(req, rec)
			if tc.name == "should return error" {
				tc.response = wrapper.ResponseError(mock.Anything, mock.Anything, c)
			} else {
				tc.response = wrapper.Response(mock.Anything, mock.Anything, http.StatusOK, c)
			}
			app.GET("/books", h.GetAllBook, nil)
			// create expectations for mocks
			uc.On("GetAllBook").
				Maybe().
				Return(tc.response)

			assert.EqualValues(t, tc.expectHttpCode, rec.Code)
		})
	}
}
