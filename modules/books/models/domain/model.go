package domain

type Book struct {
	Id     string `gorm:"column:id;primaryKey"`
	Title  string `gorm:"column:title"`
	Author string `gorm:"column:author"`
	Year   string `gorm:"column:year"`
}

func (Book) TableName() string {
	return "books"
}
