package web

import (
	"golang-api/utils/jwt"
)

type (
	RequestCreateBook struct {
		Title  string  `json:"title" validate:"required"`
		Author string  `json:"author" validate:"required"`
		Year   string  `json:"year" validate:"required"`
		Price  float64 `json:"price" validate:"required"`

		Token jwt.ClaimToken
	}

	RequestUpdateBook struct {
		Id     string  `param:"id" validate:"required,uuid4"`
		Title  string  `json:"title" validate:"required"`
		Author string  `json:"author" validate:"required"`
		Year   string  `json:"year" validate:"required"`
		Price  float64 `json:"price" validate:"required"`

		Token jwt.ClaimToken
	}

	RequestDeleteBook struct {
		Id string `param:"id" validate:"required,uuid4"`

		Token jwt.ClaimToken
	}
	RequestDetailBook struct {
		Id string `param:"id" validate:"required,uuid4"`
	}

	GetBookSheetRequest struct {
		Author string `query:"author"`
		Title  string `query:"title"`
	}
)
