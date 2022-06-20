package api

import (
	"mssgserver/api"

	"github.com/gin-gonic/gin"
)

type HostRouter struct {
}

func (r *HostRouter) InitHostRouter(Router *gin.RouterGroup) {

	Router.GET("/host/items", api.HostItems)
}
