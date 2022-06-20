package router

import (
	"github.com/gin-gonic/gin"
	"mssgserver/router/api"
)

func InitRouter(r *gin.Engine) {
	//接口
	//获取接口路由组实例
	apiRouter := api.RouterGroup{}
	ctApi := r.Group("/api")
	{
		apiRouter.InitHostRouter(ctApi)
		//apiRouter.InitBaseRouter(ctApi)
		apiRouter.InitSetRouter(ctApi)
	}

}
