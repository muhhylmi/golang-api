package web

type (
	RequestCreateBook struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Year   int64  `json:"year"`
	}

	RequestUpdateBook struct {
		Id     uint   `param:"id" json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
		Year   int64  `json:"year"`
	}

	RequestDeleteBook struct {
		Id uint `param:"id" json:"id"`
	}
	RequestDetailBook struct {
		Id uint `param:"id" json:"id"`
	}
)
