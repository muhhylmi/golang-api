package web

type (
	RequestCreateBook struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Year   int    `json:"year"`
	}

	RequestUpdateBook struct {
		Id     int    `param:"id" json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
		Year   int    `json:"year"`
	}

	RequestDeleteBook struct {
		Id int `param:"id" json:"id"`
	}
	RequestDetailBook struct {
		Id int `param:"id" json:"id"`
	}
)
