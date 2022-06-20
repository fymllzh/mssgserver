package api

import (
	"mssgserver/api"

	"github.com/gin-gonic/gin"
)

type SetRouter struct {
}

func (r *SetRouter) InitSetRouter(Router *gin.RouterGroup) {

	Router.POST("/set/delete", api.SetDelete)
	Router.POST("/set/create", api.SetItems)

}
