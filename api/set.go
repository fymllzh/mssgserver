package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetItems( ctx *gin.Context)  {
	ctx.JSON(200,gin.H{
		"do":"我是index",
	})
}
func SetDelete(  ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "我是delete",
	})
}
