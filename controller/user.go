package controller

import (
	"bookSystem/dao/mysql"
	"bookSystem/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

func RegisterHandler(c *gin.Context) {
	//1、参数验证
	p := new(model.ParamsUser)
	if err := c.ShouldBindJSON(p); err != nil {
		logger.Error(errors.New("参数绑定错误, " + err.Error()))
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//2、数据存储

	data, err := mysql.Register(p)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//3、返回信息
	c.JSON(200, gin.H{
		"msg":  "用户注册成功",
		"data": data,
	})

}

func LoginHandler(c *gin.Context) {
	p := new(model.User)
	err := c.ShouldBindJSON(p)
	if err != nil {
		logger.Error(errors.New("参数绑定错误, " + err.Error()))
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//查询数据库账号密码
	data, err := mysql.Login(p)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "登录成功",
		"data": data,
	})

}
