package web

type (
	RequestCreateBook struct {
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
		Year   string `json:"year" validate:"required"`
	}

	RequestUpdateBook struct {
		Id     string `param:"id" validate:"required,uuid4"`
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
		Year   string `json:"year" validate:"required"`
	}

	RequestDeleteBook struct {
		Id string `param:"id" validate:"required,uuid4"`
	}
	RequestDetailBook struct {
		Id string `param:"id" validate:"required,uuid4"`
	}
)
