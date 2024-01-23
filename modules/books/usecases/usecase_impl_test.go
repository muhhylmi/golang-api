package usecases

import (
	"context"
	"errors"
	"golang-api/modules/books/models/domain"
	"golang-api/modules/books/models/web"
	"golang-api/utils/config"
	"golang-api/utils/jwt"
	"golang-api/utils/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	repo_mock "golang-api/mocks/modules/books/repositories"
	gheet_mock "golang-api/mocks/utils/google_sheet"
)

func TestCreate(t *testing.T) {
	type TestCase struct {
		name            string
		findAll         []*domain.Book
		findAllErr      error
		expectStatuCode string
	}

	testCases := []TestCase{
		{
			name:            "should return error",
			findAll:         nil,
			findAllErr:      errors.New(""),
			expectStatuCode: "0101",
		},
		{
			name: "should return success",
			findAll: []*domain.Book{
				{
					Id: "1",
				},
			},
			findAllErr:      nil,
			expectStatuCode: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// create mocks
			cfg := &config.Configurations{}
			l := logger.Newlogger()
			repo := repo_mock.NewRepository(t)
			gsheet := gheet_mock.NewGoogleSheetServiceInterface(t)

			ctx := context.TODO()

			// setup expectations
			repo.On("FindAll").Maybe().Return(testCase.findAll, testCase.findAllErr)

			// call use case
			uc := NewUsecaseImpl(cfg, l, repo, gsheet, nil)
			res := uc.GetBook(ctx)

			assert.EqualValues(t, testCase.expectStatuCode, res.StatusCode)
		})
	}
}

func TestCreateBook(t *testing.T) {
	type TestCase struct {
		name            string
		createRequest   *web.RequestCreateBook
		save            *domain.Book
		saveErr         error
		expectStatuCode string
	}
	testCases := []TestCase{
		{
			name: "should return error",
			createRequest: &web.RequestCreateBook{
				Title:  "malin kundang",
				Author: "basuki",
				Year:   "2019",
				Price:  20000,
				Token: jwt.ClaimToken{
					UserId: "123",
				},
			},
			save:            nil,
			saveErr:         errors.New(""),
			expectStatuCode: "0102",
		},
		{
			name: "should return success",
			createRequest: &web.RequestCreateBook{
				Title:  "malin kundang",
				Author: "basuki",
				Year:   "2019",
				Price:  20000,
				Token: jwt.ClaimToken{
					UserId: "123",
				},
			},
			save:            &domain.Book{},
			saveErr:         nil,
			expectStatuCode: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// create mocks
			cfg := &config.Configurations{}
			l := logger.Newlogger()
			repo := repo_mock.NewRepository(t)
			gsheet := gheet_mock.NewGoogleSheetServiceInterface(t)

			ctx := context.TODO()

			// setup expectations
			repo.On("Save", mock.Anything).Maybe().Return(testCase.save, testCase.saveErr)

			// call use case
			uc := NewUsecaseImpl(cfg, l, repo, gsheet, nil)
			res := uc.CreateBook(ctx, testCase.createRequest)

			assert.EqualValues(t, testCase.expectStatuCode, res.StatusCode)
		})
	}
}
