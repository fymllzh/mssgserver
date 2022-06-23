package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"mssgserver/router/api"
)

func InitRouter(r *gin.Engine) {
	//接口
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//获取接口路由组实例
	apiRouter := api.RouterGroup{}
	ctApi := r.Group("/api")
	{
		apiRouter.InitHostRouter(ctApi)
		apiRouter.InitBaseRouter(ctApi)
		apiRouter.InitSetRouter(ctApi)
		apiRouter.InitLoginRouter(ctApi)
	}

}
