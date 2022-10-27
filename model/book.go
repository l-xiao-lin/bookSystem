package model

type Book struct {
	ID    int     `gorm:"primaryKey" json:"id"`
	Name  string  `gorm:"not null" binding:"required" json:"name"`
	Desc  string  `gorm:"not null" binding:"required" json:"desc"`
	Users []*User `gorm:"many2many:book_user"`
}

func (Book) TableName() string {
	return "book"

}
