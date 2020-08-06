package main

import (
	"BigStrawberryPlanet/web/controller"
	"BigStrawberryPlanet/web/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func main() {

	engine := gin.New()
	engine.Use(middleware.LoggerToFile(Logger))

	engine.GET("/", func(c * gin.Context) {
		c.Writer.WriteString("hello big strawberry planet...")

	})

	// 根据服务划分路由
	user := engine.Group("/user")
	{
		user.GET("/get", controller.UserGet)
		user.POST("/post", controller.UserPost)
		user.PUT("/put", controller.UserPut)
		user.DELETE("/delete", controller.UserDelete)
	}

	engine.Run()

}
