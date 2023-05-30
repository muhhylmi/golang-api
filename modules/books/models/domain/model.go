package domain

type Book struct {
	Id     uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Title  string `gorm:"column:tittle"`
	Author string `gorm:"column:author"`
	Year   int64  `gorm:"column:year"`
}

func (Book) TableName() string {
	return "books"
}
