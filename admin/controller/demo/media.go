package demo

import (
	"rtsys/admin/lib"
	. "rtsys/common/models/demo"

	"github.com/gin-gonic/gin"
)

type MediaController struct {
	lib.Controller
}

func (c *MediaController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.POST(c.Dispatch("index", false, c.FindList))
	group.GET(c.Dispatch("add", false, c.Add))
	group.POST(c.Dispatch("add", false, c.AddSave))
	group.GET(c.Dispatch("edit", false, c.Edit))
	group.POST(c.Dispatch("edit", false, c.EditSave))
	group.POST(c.Dispatch("drop", false, c.Drop))
	group.POST(c.Dispatch("drop_all", false, c.DropAll))
}

func NewMediaController() *MediaController {
	c := &MediaController{}
	c.Id = "media"
	c.Group = "demo"

	c.IModel = NewMediaModel()
	return c
}

func (c *MediaController) FindList(ctx *gin.Context) {
	c.GetIndex(ctx, NewMediaModel().NewSlice(), "", nil, nil)
}

func (c *MediaController) Add(ctx *gin.Context) {
	c.Render(ctx, "add", nil)
}
