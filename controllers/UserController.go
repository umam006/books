package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

func (ctl UserController) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
