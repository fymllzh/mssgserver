package api

import (
	"mssgserver/api"

	"github.com/gin-gonic/gin"
)

type SetRouter struct {
}

func (r *SetRouter) InitSetRouter(Router *gin.RouterGroup) {

	Router.GET("/set/delete", api.SetDelete)
	Router.GET("/set/create", api.SetItems)

}
