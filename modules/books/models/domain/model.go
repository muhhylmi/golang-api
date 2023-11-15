package domain

type Book struct {
	Id        string `gorm:"column:id;primaryKey"`
	Title     string `gorm:"column:title"`
	Author    string `gorm:"column:author"`
	Year      string `gorm:"column:year"`
	CreatedAt int64  `gorm:"column:created_at"`
	CreatedBy string `gorm:"column:created_by"`
	UpdatedAt int64  `gorm:"column:updated_at"`
	UpdatedBy string `gorm:"column:updated_by"`
}

func (Book) TableName() string {
	return "books"
}
