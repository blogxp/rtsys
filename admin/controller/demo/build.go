package demo

import (
	"rtsys/admin/lib"

	"github.com/gin-gonic/gin"
)

func (c *BuildController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
}

type BuildController struct {
	lib.Controller
}

func NewBuildController() *BuildController {
	c := &BuildController{}
	c.Id = "build"
	c.Group = "demo"
	return c
}

func (c *BuildController) Index(ctx *gin.Context) {
	c.Render(ctx, "index", nil)
}
