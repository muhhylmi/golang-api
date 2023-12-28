package web

type ResponseSheetBook struct {
	BookName string  `json:"bookName"`
	Author   string  `json:"author"`
	Year     string  `json:"year"`
	Total    float64 `json:"totalHarga"`
}
