package model

type BookUser struct {
	UserID int `gorm:"primaryKey"`
	BookID int `gorm:"primaryKey"`
}

func (BookUser) TableName() string {
	return "book_user"
}
