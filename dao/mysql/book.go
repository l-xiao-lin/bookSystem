package mysql

import (
	"bookSystem/model"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
)

func CreateBook(data *model.ParamsCreateBook) (err error) {
	//1、判断book表中是否存在同样的图书
	row := DB.Where("name = ?", data.Name).Find(&model.Book{}).RowsAffected
	if row == 1 {
		logger.Error(errors.New("图书信息已存在"))
		return errors.New("图书信息已存在")
	}

	//2、查询user表中的数据
	var users []*model.User
	DB.Debug().Where("username LIKE ?", "%huawei%").Find(&users)
	fmt.Println(users)

	//3、将新增的书本与用户绑定
	book := model.Book{
		Name:  data.Name,
		Desc:  data.Desc,
		Users: users,
	}
	tx := DB.Create(&book)
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		logger.Error(errors.New("创建图书错误"))
		return errors.New("创建图书错误")
	}
	return nil
}

func GetBookList() (bookList *model.RespBookList, err error) {
	var books []model.Book
	tx := DB.Preload("Users").Find(&books)
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		logger.Error(errors.New("获取图书列表失败"))
		return nil, errors.New("获取图书列表失败")
	}
	total := len(books)
	bookList = &model.RespBookList{
		Items: books,
		Total: total,
	}
	return bookList, nil
}

func GetBookDetail(bookId int) (*model.Book, error) {
	book := &model.Book{ID: int(bookId)}
	tx := DB.Preload("Users").Find(book)
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		logger.Error(errors.New("获取图书详情失败"))
		return nil, errors.New("获取图书详情失败")
	}
	return book, nil
}

func EditBook(data *model.Book) (*model.Book, error) {
	//1、先查询旧的数据
	oldBook := model.Book{ID: data.ID}
	tx := DB.Model(&oldBook).Updates(model.Book{Name: data.Name, Desc: data.Desc})
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		logger.Error(errors.New("更新图书信息失败"))
		return nil, errors.New("更新图书信息失败")
	}
	newBook := &model.Book{ID: data.ID}
	DB.Find(newBook)
	return newBook, nil
}

func DeleteBook(id int) (err error) {

	tx := DB.Select("Users").Delete(&model.Book{ID: id})
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		logger.Error(errors.New("删除图书失败"))
		return errors.New("删除图书失败")
	}
	return nil

}
