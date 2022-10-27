package controller

import (
	"bookSystem/dao/mysql"
	"bookSystem/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"strconv"
)

func CreateBookHandler(c *gin.Context) {
	p := new(model.ParamsCreateBook)
	if err := c.ShouldBindJSON(p); err != nil {
		logger.Error(errors.New("参数绑定失败, " + err.Error()))
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//调用图书数据存储
	if err := mysql.CreateBook(p); err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "成功创建图书",
	})

}

func GetBookListHandler(c *gin.Context) {
	data, err := mysql.GetBookList()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "成功获取图书列表",
		"data": data,
	})
}

func GetBookDetailHandler(c *gin.Context) {
	strID := c.Param("id")
	//bookId, _ := strconv.ParseInt(strID, 10, 64)
	bookId, _ := strconv.Atoi(strID)
	data, err := mysql.GetBookDetail(bookId)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "成功获取图书详情",
		"data": data,
	})
}

func EditBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBindJSON(p); err != nil {
		logger.Error(errors.New("参数绑定错误, " + err.Error()))
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//调用修改图书
	data, err := mysql.EditBook(p)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "成功修改图书信息",
		"data": data,
	})

}

func DeleteBookHandler(c *gin.Context) {
	strId := c.Param("id")
	bookId, _ := strconv.Atoi(strId)

	//调用数据删除操作
	err := mysql.DeleteBook(bookId)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "图书信息已删除",
	})

}
