// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package common

import (
	"html/template"
	"rtsys/admin/lib"
	"rtsys/admin/services"
	"rtsys/common/services/system"
	"rtsys/utils/types"

	"github.com/gin-gonic/gin"
)

func (c *IndexController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.GET(c.Dispatch("home", false, c.Home))
}

type IndexController struct {
	lib.Controller
}

// NewIndexController 创建控制并初始化参数
func NewIndexController() *IndexController {
	c := &IndexController{}
	c.Group = ""
	c.Id = "index"
	return c
}

func (c *IndexController) Index(ctx *gin.Context) {
	userInfo := map[string]string{"nick_name": "", "struct": ""}

	loginData := services.GetLoginByCtx(ctx)
	if loginData != nil {
		userInfo["nick_name"] = loginData.NickName
		userInfo["struct"] = system.NewStructService().GetName(loginData.StructId, false)
	}

	list := services.AdminMenuShowList(ctx)
	html := ""
	if list != nil && len(list) > 0 {
		html = services.AdminMenuToHtml(list, 0)
	}
	c.Render(ctx, "index", gin.H{"user": userInfo, "menuHtml": &types.HtmlShow{Html: template.HTML(html)}})
}

func (c *IndexController) Home(ctx *gin.Context) {
	c.Render(ctx, "home", nil)
}
