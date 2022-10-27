package routers

import (
	"bookSystem/controller"
	"bookSystem/middleware"
	"github.com/gin-gonic/gin"
)

var Router router

type router struct {
}

func (r *router) InitSetupRouter(router *gin.Engine) {

	router.GET("/demo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "bookSystem...",
		})
	})

	router.POST("/register", controller.RegisterHandler)
	router.POST("/login", controller.LoginHandler)

	v1 := router.Group("/api/v1/")
	v1.Use(middleware.AuthMiddleware())

	v1.POST("/book", controller.CreateBookHandler)
	v1.GET("/book", controller.GetBookListHandler)
	v1.GET("/book/:id", controller.GetBookDetailHandler)
	v1.PUT("/book/", controller.EditBookHandler)
	v1.DELETE("/book/:id", controller.DeleteBookHandler)

}
