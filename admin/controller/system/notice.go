// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"rtsys/admin/lib"
	. "rtsys/common/models/system"
	"rtsys/common/services/system"
	"rtsys/utils/core"

	"github.com/gin-gonic/gin"
)

// Route 定义该控制器的路由
func (c *NoticeController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.POST(c.Dispatch("index", false, c.FindList))
	group.GET(c.Dispatch("add", false, c.Add))
	group.POST(c.Dispatch("add", false, c.AddSave))
	group.GET(c.Dispatch("edit", false, c.Edit))
	group.POST(c.Dispatch("edit", false, c.EditSave))
	group.POST(c.Dispatch("drop", false, c.Drop))
}

type NoticeController struct {
	lib.Controller
}

// NewNoticeController 创建控制并初始化参数
func NewNoticeController() *NoticeController {
	c := &NoticeController{}
	c.Group = "system"
	c.Id = "notice"
	c.IModel = NewNoticeModel()
	c.HookAction.EditRender = c.editRender
	return c
}
func (c *NoticeController) FindList(ctx *gin.Context) {
	c.GetIndex(ctx, NewNoticeModel().NewSlice(), "", nil, nil)
}

func (c *NoticeController) Add(ctx *gin.Context) {
	c.Render(ctx, "add", gin.H{"typeList": system.NewNoticeService().TypeList()})
}
func (c *NoticeController) editRender(ctx *gin.Context, model core.IModel) {
	c.Render(ctx, "edit", gin.H{"typeList": system.NewNoticeService().TypeList(), "info": model})
}
