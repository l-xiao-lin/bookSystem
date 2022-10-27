package model

import (
	"time"
)

type ParamsUser struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
	CreateTime time.Time
}

type ParamsCreateBook struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type RespLogin struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type RespBookList struct {
	Items []Book
	Total int
}
