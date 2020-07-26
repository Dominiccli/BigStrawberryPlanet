package main

import (
	"BigStrawberryPlanet/web/controller"
	"github.com/gin-gonic/gin"
)


// 添加 gin 框架开发

func main() {

	router := gin.Default()
	router.GET("/", func(c * gin.Context) {
		c.Writer.WriteString("hello big strawberry planet...")
	})

	router.GET("/userGet", controller.UserGet)
	router.POST("/userPost", controller.UserPost)
	router.PUT("/userPut", controller.UserPut)
	router.DELETE("/userDelete", controller.UserDelete)

	router.Run()

}
