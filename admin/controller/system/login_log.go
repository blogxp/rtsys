// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"rtsys/admin/lib"
	"rtsys/common/daos/system"
	. "rtsys/common/models/system"

	"github.com/gin-gonic/gin"
)

// Route 定义该控制器的路由
func (c *LoginLogController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.POST(c.Dispatch("index", false, c.FindList))
	group.POST(c.Dispatch("trash", false, c.Trash))
	group.POST(c.Dispatch("drop_all", false, c.DropAll))
}

type LoginLogController struct {
	lib.Controller
}

// NewLoginLogController 创建控制并初始化参数
func NewLoginLogController() *LoginLogController {
	c := &LoginLogController{}
	c.Id = "login_log"
	c.Group = "system"
	c.IModel = NewLoginLogModel()
	return c
}

// FindList 获取列表json
func (c *LoginLogController) FindList(ctx *gin.Context) {
	c.GetIndex(ctx, NewLoginLogModel().NewSlice(), "", nil, nil)
}

func (c *LoginLogController) Trash(ctx *gin.Context) {
	err := system.NewLoginLogDao().Trash()
	if err != nil {
		c.Error(ctx, "操作失败："+err.Error())
		return
	}
	c.Success(ctx, "操作成功")
}
