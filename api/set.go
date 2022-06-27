package api

import (
	"github.com/gin-gonic/gin"
	"mssgserver/utils"
	"net/http"
	"github.com/gin-contrib/sessions"
)

func SetItems( c *gin.Context)  {
    s:= sessions.Default(c)
    admin_name := s.Get("admin_name")
	cttask_token := s.Get("cttask_token")
	utils.Logger.Info("login success")

	c.JSON(200,gin.H{
		"do":"登陆成功",
		"admin_name":admin_name,
		"cttask_token":cttask_token,
	})
}
func SetDelete(  c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "我是delete",
	})
}
