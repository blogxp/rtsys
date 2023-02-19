// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"errors"
	"rtsys/admin/lib"
	. "rtsys/common/daos/system"
	. "rtsys/common/models/system"
	. "rtsys/common/services/system"
	"rtsys/utils/core"
	"rtsys/utils/types"

	"github.com/gin-gonic/gin"
)

// Route 定义该控制器的路由
func (c *MenuController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.POST(c.Dispatch("index", false, c.FindList))
	group.GET(c.Dispatch("add", false, c.Add))
	group.POST(c.Dispatch("add", false, c.AddSave))
	group.GET(c.Dispatch("edit", false, c.Edit))
	group.POST(c.Dispatch("edit", false, c.EditSave))
	group.POST(c.Dispatch("drop", false, c.Drop))
	group.GET(c.Dispatch("tree", false, c.Tree))
	group.POST(c.Dispatch("tree", false, c.TreeList))
}

type MenuController struct {
	lib.Controller
}

// NewMenuController 创建控制并初始化参数
func NewMenuController() *MenuController {
	c := &MenuController{}
	c.Group = "system"
	c.Id = "menu"
	c.IModel = NewMenuModel()

	c.HookAction.EditRender = c.editRender
	c.HookAction.SaveBefore = c.saveBefore
	c.HookAction.DropBefore = c.dropBefore
	c.HookAction.DropAfter = c.dropAfter
	return c
}
func (c *MenuController) FindList(ctx *gin.Context) {
	c.GetIndex(ctx, NewMenuModel().NewSlice(), "", nil, []types.KeyVal{{Key: "parent_id"}, {Key: "list_sort"}})
}

func (c *MenuController) Add(ctx *gin.Context) {
	c.Render(ctx, "add", gin.H{"typeList": NewMenuService().TypeList()}, c.Group+"/menu/icon")
}

func (c *MenuController) editRender(ctx *gin.Context, model core.IModel) {
	info := model.(*MenuModel)
	parentName := NewMenuService().GetParentName(info.ParentId)
	c.Render(ctx, "edit", gin.H{"info": model, "typeList": NewMenuService().TypeList(), "parentName": parentName}, c.Group+"/menu/icon")
}

func (c *MenuController) saveBefore(ctx *gin.Context, model core.IModel, operate string) error {
	if operate != "edit" {
		return nil
	}
	info := model.(*MenuModel)
	if info.ParentId == info.Id {
		return errors.New("所属菜单不能是自己")
	}
	return nil
}

func (c *MenuController) dropBefore(ctx *gin.Context, model core.IModel) error {
	info, err := core.NewDao(c.IModel).GetInfoByField("parent_id", model.GetId(), "id")
	if err != nil || info.GetId() == "" {
		return nil
	}
	return errors.New("有子菜单，无法删除")
}

func (c *MenuController) dropAfter(ctx *gin.Context, model core.IModel) {
	_ = NewRoleMenuDao().DeleteByMenuId(model.GetId())
}

func (c *MenuController) Tree(ctx *gin.Context) {
	root := ctx.DefaultQuery("root", "0")
	id := ctx.DefaultQuery("id", "0")
	c.Render(ctx, "tree", gin.H{"id": id, "root": root})
}
func (c *MenuController) TreeList(ctx *gin.Context) {
	root := ctx.DefaultQuery("root", "0")
	list := NewMenuService().TreeList(root)
	c.Success(ctx, "", list)
}
