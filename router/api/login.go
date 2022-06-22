package api

import (
	"mssgserver/api"

	"github.com/gin-gonic/gin"
)

type LoginRouter struct {
}

func (r *HostRouter) InitLoginRouter(Router *gin.RouterGroup) {

	Router.POST("/login", api.LoginAuth)
}
