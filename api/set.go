package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mssgserver/utils"
	"net/http"
)

func SetItems( c *gin.Context)  {

	defer func() {
		if err := recover(); err != nil {
			utils.Logger.Info(err)
			if str, ok := err.(string); ok {
				utils.Logger.Info("set异常是 " + str)
			}
		}
	}()

    s:= sessions.Default(c)
    admin_name := s.Get("admin_name")
	cttask_token := s.Get("cttask_token")
	utils.Logger.Info("login success")

	//返回用户数据
	var user Logins
		data, err := user.Items()
		if err != nil {
			panic(err)
		}

	c.JSON(200,gin.H{
		"do":"登陆成功",
		"admin_name":admin_name,
		"cttask_token":cttask_token,
		"data":data,
	})
}
func SetDelete(  c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "我是delete",
	})
}
