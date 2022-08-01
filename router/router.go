package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"mssgserver/router/api"
	"mssgserver/server"
	"mssgserver/utils"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	//接口
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	//这样些会报import cycle not allowed 不知道怎么解决
	//r.Use(utils.IpWhiteList())
	//ip 白名单
	r.Use(func(ctx *gin.Context) {
		ip := utils.GetRequestIP(ctx)
		utils.Logger.Info("ip_request:", ip)
		var ipwhite server.Ip
		ipinfo, _ := ipwhite.Allow(ip)

		if ipinfo.Id <= 0 {
			utils.Logger.Info("ip_denny:", ip)
			ctx.JSON(http.StatusOK, gin.H{
				"errno":   403,
				"message": "Access Denny.",
				"ip":      ip,
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	})

	//获取接口路由组实例
	apiRouter := api.RouterGroup{}
	ctApi := r.Group("/api")
	//加入统计时间的中间件
	ctApi.Use(utils.StatCost())
	{
		apiRouter.InitHostRouter(ctApi)
		apiRouter.InitBaseRouter(ctApi)
		apiRouter.InitSetRouter(ctApi)
		apiRouter.InitLoginRouter(ctApi)
	}

}
