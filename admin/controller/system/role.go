// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"errors"
	"rtsys/admin/lib"
	"rtsys/admin/services"
	. "rtsys/common/daos/system"
	. "rtsys/common/models/system"
	. "rtsys/common/services/system"
	"rtsys/utils/core"
	"strings"

	"github.com/gin-gonic/gin"
)

// Route 定义该控制器的路由
func (c *RoleController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.POST(c.Dispatch("index", false, c.FindList))
	group.GET(c.Dispatch("add", false, c.Add))
	group.POST(c.Dispatch("add", false, c.AddSave))
	group.GET(c.Dispatch("edit", false, c.Edit))
	group.POST(c.Dispatch("edit", false, c.EditSave))
	group.POST(c.Dispatch("drop", false, c.Drop))
	group.POST(c.Dispatch("drop_all", false, c.DropAll))
	group.POST(c.Dispatch("set_status", false, c.SetStatus))
	group.GET(c.Dispatch("auth", false, c.Auth))
	group.POST(c.Dispatch("auth", false, c.AuthSave))
	group.GET(c.Dispatch("data_scope", false, c.DataScope))
	group.POST(c.Dispatch("data_scope", false, c.DataScopeSave))
}

type RoleController struct {
	lib.Controller
}

// NewRoleController 创建控制并初始化参数
func NewRoleController() *RoleController {
	c := &RoleController{}
	c.Id = "role"
	c.Group = "system"
	c.IModel = NewRoleModel()

	c.HookAction.SaveBefore = c.saveBefore
	c.HookAction.DropAfter = c.dropAfter
	return c
}

func (c *RoleController) Index(ctx *gin.Context) {
	c.Render(ctx, "index", gin.H{"root_id": core.G_ROLE_ID})
}

// FindList 获取列表json
func (c *RoleController) FindList(ctx *gin.Context) {
	c.GetIndex(ctx, NewRoleModel().NewSlice(), "", nil, nil)
}

// saveBefore 保存前验证编码唯一性
func (c *RoleController) saveBefore(ctx *gin.Context, model core.IModel, operate string) error {
	if operate == "add" || operate == "edit" {
		model := model.(*RoleModel)
		info, err := core.NewDao(c.IModel).GetInfoByField("role_key", model.RoleKey, "")
		if err != nil {
			return nil
		}
		if (operate == "edit" && info.GetId() != model.GetId()) || operate == "add" {
			return errors.New("角色标识已存在")
		}
	}
	return nil
}

// 删除后操作
func (c *RoleController) dropAfter(ctx *gin.Context, model core.IModel) {
	//删除菜单和数据授权
	_ = NewRoleStructDao().DeleteByRoleId(model.GetId()) //角色数据权限组织
	_ = NewRoleMenuDao().DeleteByRoleId(model.GetId())   //角色菜单信息
	_ = NewAdminRoleDao().DeleteByRoleId(model.GetId())  //角色的管理员信息
}

// Auth 角色授权菜单页面
func (c *RoleController) Auth(ctx *gin.Context) {
	roleId := ctx.DefaultQuery("role_id", "")
	if roleId == "" {
		c.Jump(ctx, errors.New("参数错误"))
		return
	}
	model := NewRoleDao().GetInfoById(roleId)
	if model == nil {
		c.Jump(ctx, errors.New("信息查询错误"))
		return
	}
	menuList := NewRoleMenuService().GetRoleMenuIdList(model.Id)
	c.Render(ctx, "auth", gin.H{"id": model.Id, "name": model.Name, "menuList": strings.Join(menuList, ",")})
}

func (c *RoleController) AuthSave(ctx *gin.Context) {
	form := &lib.RoleMenuDataAuth{}
	err := ctx.ShouldBind(form)
	if err != nil {
		c.Error(ctx, "数据绑定失败")
		return
	}
	err1 := NewRoleService().SaveMenuAuth(form)
	if err1 != nil {
		c.Error(ctx, err1.Error())
		return
	}
	c.Success(ctx, "菜单授权成功")
}

func (c *RoleController) DataScope(ctx *gin.Context) {
	roleId := ctx.DefaultQuery("role_id", "")
	if roleId == "" {
		c.Jump(ctx, errors.New("参数错误"))
		return
	}
	model := NewRoleDao().GetInfoById(roleId)

	if model == nil {
		c.Jump(ctx, errors.New("信息查询错误"))
		return
	}
	structList := NewRoleStructService().GetRoleStructIdList(model.Id)
	c.Render(ctx, "datascope", gin.H{"typeList": services.DataScopeTypeList(), "info": model, "structList": strings.Join(structList, ",")})
}

func (c *RoleController) DataScopeSave(ctx *gin.Context) {
	form := &lib.RoleMenuDataAuth{}
	err := ctx.ShouldBind(form)
	if err != nil {
		c.Error(ctx, "数据绑定失败")
		return
	}
	err1 := NewRoleService().SaveDataScopeAuth(form)
	if err1 != nil {
		c.Error(ctx, err1.Error())
		return
	}
	c.Success(ctx, "数据权限授权成功")
}
