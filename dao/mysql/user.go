package mysql

import (
	"bookSystem/model"
	"bookSystem/pkg"
	"errors"
	"github.com/wonderivan/logger"
	"time"
)

func Register(data *model.ParamsUser) (user *model.User, err error) {

	//1、判断用户是否存在
	row := DB.Where("username = ?", data.Username).Find(&model.User{}).RowsAffected
	if row == 1 {
		logger.Error(errors.New("该用户已存在"))
		return nil, errors.New("该用户已存在")
	}

	//2、判断两次密码是否一致
	if data.Password != data.RePassword {
		logger.Error(errors.New("两次输入的密码不一致,请重新输入"))
		return nil, errors.New("两次输入的密码不一致,请重新输入")
	}

	//3、密码进行加密
	md5Password := pkg.GetMd5String(data.Password)

	user = &model.User{
		Username:   data.Username,
		Password:   md5Password,
		CreateTime: time.Now(),
	}

	tx := DB.Create(user)
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		logger.Error(errors.New("用户注册失败," + err.Error()))
		return nil, errors.New("用户注册失败," + err.Error())
	}
	return user, nil

}

func Login(data *model.User) (user *model.RespLogin, err error) {
	//将用户传进来的密码进行md5加密，然后再查询数据库
	md5Password := pkg.GetMd5String(data.Password)

	u1 := model.User{}
	row := DB.Where("username = ? AND password = ?", data.Username, md5Password).Find(&u1).RowsAffected
	if row == 0 {
		logger.Error(errors.New("用户名或密码错误"))
		return nil, errors.New("用户名或密码错误")
	}

	//生成token
	token, err := pkg.GenToken(u1.Username)
	if err != nil {
		return nil, err
	}

	result := &model.RespLogin{
		Username: u1.Username,
		Token:    token,
	}
	return result, nil

}
