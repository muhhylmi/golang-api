package domain

type Users struct {
	Id       string `gorm:"column:id;primaryKey"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Gender   string `gorm:"column:gender"`
}

func (Users) TableName() string {
	return "users"
}
