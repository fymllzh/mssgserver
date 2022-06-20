package api

import "github.com/gin-gonic/gin"

func HostItems( ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"mes":"w shi  host items",
	})
}
