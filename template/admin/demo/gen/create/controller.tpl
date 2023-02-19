// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package __GROUP__

import (
	"rtsys/admin/lib"
	__MODEL_PACKAGE__
	"github.com/gin-gonic/gin"
)

// Route 定义该控制器的路由
func (c *__CONTROLLER__) Route(engine *gin.Engine,group *gin.RouterGroup) {
    group.GET(c.Dispatch("index", false, c.Index))
    group.POST(c.Dispatch("index", false, c.FindList))
    group.GET(c.Dispatch("add", false, c.Add))
    group.POST(c.Dispatch("add", false, c.AddSave))
    group.GET(c.Dispatch("edit", false, c.Edit))
    group.POST(c.Dispatch("edit", false, c.EditSave))
    group.POST(c.Dispatch("drop", false, c.Drop))
    group.POST(c.Dispatch("drop_all", false, c.DropAll))
}

type __CONTROLLER__ struct {
	lib.Controller
}

// New__CONTROLLER__ 创建控制并初始化参数
func New__CONTROLLER__() *__CONTROLLER__ {
	c := &__CONTROLLER__{}
	c.Id = "__ID__"
	c.Group = "__GROUP__"
	__MODEL_NEW__
	return c
}

// FindList 获取列表json
func (c *__CONTROLLER__) FindList(ctx *gin.Context) {
	c.GetIndex(ctx, __MODEL_SLICE__, "", nil, nil)
}
