package controller

import "github.com/gin-gonic/gin"



func UserGet(ctx * gin.Context) {
	ctx.Writer.WriteString("user get ...")
}

func UserPost(ctx * gin.Context) {
	ctx.Writer.WriteString("user get ...")
}

func UserPut(ctx * gin.Context) {
	ctx.Writer.WriteString("user put ...")
}

func UserDelete(ctx * gin.Context) {
	ctx.Writer.WriteString("user delete ...")
}

