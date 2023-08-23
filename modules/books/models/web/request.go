package web

type (
	RequestCreateBook struct {
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
		Year   int64  `json:"year" validate:"required"`
	}

	RequestUpdateBook struct {
		Id     uint   `param:"id" validate:"required"`
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
		Year   int64  `json:"year" validate:"required"`
	}

	RequestDeleteBook struct {
		Id uint `param:"id" json:"id"`
	}
	RequestDetailBook struct {
		Id uint `param:"id" json:"id"`
	}
)
