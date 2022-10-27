package model

import "time"

type User struct {
	ID         int    `gorm:"primaryKey"`
	Username   string `gorm:"not null" binding:"required" json:"username"`
	Password   string `gorm:"not null" binding:"required" json:"password"`
	CreateTime time.Time
}

func (User) TableName() string {
	return "user"
}
