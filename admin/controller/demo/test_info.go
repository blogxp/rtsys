// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package demo

import (
	"errors"
	"rtsys/admin/lib"
	"rtsys/admin/services"
	. "rtsys/common/models/demo"
	"rtsys/common/services/system"
	"rtsys/utils/core"
	"rtsys/utils/trans"

	"github.com/gin-gonic/gin"
)

// Route 定义该控制器的路由
func (c *TestInfoController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.POST(c.Dispatch("index", false, c.FindList))
	group.GET(c.Dispatch("add", false, c.Add))
	group.POST(c.Dispatch("add", false, c.AddSave))
	group.GET(c.Dispatch("edit", false, c.Edit))
	group.POST(c.Dispatch("edit", false, c.EditSave))
	group.POST(c.Dispatch("drop", false, c.Drop))
	group.POST(c.Dispatch("drop_all", false, c.DropAll))
}

type TestInfoController struct {
	lib.Controller
}

// NewTestInfoController 创建控制并初始化参数
func NewTestInfoController() *TestInfoController {
	c := &TestInfoController{}
	c.Id = "test_info"
	c.Group = "demo"
	c.IModel = NewTestInfoModel()

	c.HookAction.IndexAfter = c.indexAfter
	c.HookAction.EditRender = c.editRender
	c.HookAction.SaveBefore = c.saveBefore
	c.HookAction.DropBefore = c.dropBefore
	return c
}

// FindList 获取列表json
func (c *TestInfoController) FindList(ctx *gin.Context) {
	//获取数据权限的where语句和占位替换args，两种方法，第二种必须存在一个struct_levels字段
	//params, args := services.NewDataScopeFilterByCtx(ctx).GetQueryParams("struct_id", "user_id")
	params, args := services.NewDataScopeFilterByCtx(ctx).GetQueryParamsByLevel("struct_id", "struct_levels", "user_id")
	c.GetIndex(ctx, NewTestInfoModel().NewSlice(), params, args, nil)
}

func (c *TestInfoController) indexAfter(ctx *gin.Context, list []any) []any {
	reList := make([]any, 0)
	for _, item := range list {
		info := item.(TestInfoModel)
		model := trans.StructToMap(info, true)
		model["struct_name"] = system.NewStructService().GetName(info.StructId, false)
		model["user_name"] = system.NewAdminService().GetName(info.UserId)
		reList = append(reList, model)
	}
	return reList
}

func (c *TestInfoController) editRender(ctx *gin.Context, model core.IModel) {
	info := model.(*TestInfoModel)
	dataScope := services.NewDataScopeFilterByCtx(ctx)
	if !dataScope.CheckByFiledLevels(info.StructId, info.StructLevels, info.UserId) {
		c.Jump(ctx, errors.New("无数据权限"))
		return
	}
	c.Render(ctx, "edit", gin.H{"info": model})
}

func (c *TestInfoController) saveBefore(ctx *gin.Context, model core.IModel, operate string) error {
	if operate == "add" || operate == "edit" {
		info := model.(*TestInfoModel)
		if operate == "edit" {
			dataScope := services.NewDataScopeFilterByCtx(ctx)
			if filed := dataScope.CheckByFiled(info.StructId, info.UserId); !filed {
				return errors.New("无权限此操作")
			}
		}

		loginData := services.GetLoginByCtx(ctx)
		info.StructId = loginData.StructId
		info.UserId = loginData.Id

		//组织信息
		structInfo := system.NewStructService().GetInfo(loginData.StructId)
		if structInfo == nil {
			return errors.New("组织信息查询失败")
		}
		info.StructLevels = structInfo.Levels
	}
	return nil
}

func (c *TestInfoController) dropBefore(ctx *gin.Context, model core.IModel) error {
	info := model.(*TestInfoModel)
	dataScope := services.NewDataScopeFilterByCtx(ctx)
	if filed := dataScope.CheckByFiled(info.StructId, info.UserId); !filed {
		return errors.New("无权限此操作")
	}
	return nil
}
