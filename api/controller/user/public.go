package user

import (
	"rtsys/api/lib"

	"github.com/gin-gonic/gin"
)

type PublicApi struct {
	lib.BaseApi
}

func NewPublicApi() *PublicApi {
	c := &PublicApi{}
	c.Id = "public"
	return c
}

func (c *PublicApi) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", c.Index))
}

func (c *PublicApi) Index(ctx *gin.Context) {
	c.Error(ctx)
}
