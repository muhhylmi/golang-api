package web

import "golang-api/utils"

type (
	RequestCreateBook struct {
		Title  string  `json:"title" validate:"required"`
		Author string  `json:"author" validate:"required"`
		Year   string  `json:"year" validate:"required"`
		Price  float64 `json:"price" validate:"required"`

		Token utils.ClaimToken
	}

	RequestUpdateBook struct {
		Id     string  `param:"id" validate:"required,uuid4"`
		Title  string  `json:"title" validate:"required"`
		Author string  `json:"author" validate:"required"`
		Year   string  `json:"year" validate:"required"`
		Price  float64 `json:"price" validate:"required"`

		Token utils.ClaimToken
	}

	RequestDeleteBook struct {
		Id string `param:"id" validate:"required,uuid4"`

		Token utils.ClaimToken
	}
	RequestDetailBook struct {
		Id string `param:"id" validate:"required,uuid4"`
	}
)
