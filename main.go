package main

import (
	"bookSystem/dao/mysql"
	"bookSystem/routers"
	"bookSystem/settings"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//初始化配置文件
	if err := settings.InitConf(); err != nil {
		fmt.Println(err.Error())
		return
	}

	//加载路由
	routers.Router.InitSetupRouter(r)

	//初始化mysql
	if err := mysql.InitMysqlDB(settings.Conf.MysqlConf); err != nil {
		fmt.Println(err.Error())
		return
	}

	r.Run(":8080")

}
