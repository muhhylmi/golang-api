package domain

import (
	domainBooks "golang-api/modules/books/models/domain"
	"golang-api/modules/users/models/domain"
)

type Cart struct {
	Id        string  `gorm:"column:id;primaryKey"`
	UserId    string  `gorm:"column:user_id"`
	Price     float64 `gorm:"column:price"`
	Status    string  `gorm:"column:status"`
	CreatedBy string  `gorm:"created_by"`
	UpdatedBy string  `gorm:"updated_by"`
	CreatedAt int64   `gorm:"autoCreateTime;->;<-:create;" json:"createdAt"`
	UpdatedAt int64   `gorm:"autoUpdateTime" json:"updatedAt"`

	Details []CartDetail `gorm:"foreignKey:CartId;references:Id"`
	User    domain.Users `gorm:"foreignKey:UserId;references:Id"`
}

func (Cart) TableName() string {
	return "carts"
}

type CartDetail struct {
	Id        string `gorm:"column:id;primaryKey"`
	CartId    string `gorm:"column:cart_id"`
	BookId    string `gorm:"column:book_id"`
	Qty       int    `gorm:"column:qty"`
	CreatedBy string `gorm:"created_by"`
	UpdatedBy string `gorm:"updated_by"`
	CreatedAt int64  `gorm:"autoCreateTime;->;<-:create;" json:"createdAt"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updatedAt"`

	Books domainBooks.Book `gorm:"foreignKey:BookId;references:Id"`
}

func (CartDetail) TableName() string {
	return "cart_details"
}
