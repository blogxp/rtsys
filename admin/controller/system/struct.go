// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"errors"
	"fmt"
	"rtsys/admin/lib"
	"rtsys/admin/services"
	. "rtsys/common/daos/system"
	. "rtsys/common/models/system"
	. "rtsys/common/services/system"
	"rtsys/utils/core"
	"rtsys/utils/trans"
	"rtsys/utils/types"
	"strings"

	"github.com/gin-gonic/gin"
)

// Route 定义该控制器的路由
func (c *StructController) Route(engine *gin.Engine, group *gin.RouterGroup) {
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

type StructController struct {
	lib.Controller
}

// NewStructController 创建控制并初始化参数
func NewStructController() *StructController {
	c := &StructController{}
	c.Group = "system"
	c.Id = "struct"
	c.IModel = NewStructModel()

	c.HookAction.IndexAfter = c.indexAfter
	c.HookAction.EditRender = c.editRender
	c.HookAction.SaveBefore = c.saveBefore
	c.HookAction.SaveAfter = c.saveAfter
	c.HookAction.DropBefore = c.dropBefore
	c.HookAction.DropAfter = c.dropAfter
	return c
}

func (c *StructController) FindList(ctx *gin.Context) {
	c.GetIndex(ctx, NewStructModel().NewSlice(), "", nil, []types.KeyVal{{Key: "parent_id"}, {Key: "list_sort"}})
}

func (c *StructController) indexAfter(ctx *gin.Context, list []any) []any {
	reList := make([]any, len(list))
	for index, item := range list {
		info := item.(StructModel)
		model := trans.StructToMap(info, true)
		model["type_name"] = NewConfigService().GetStructTypeName(info.Type)
		reList[index] = model
	}
	return reList
}

func (c *StructController) Add(ctx *gin.Context) {
	rootId := core.G_STRUCT_ID
	rootName := NewStructService().GetName(rootId, false)
	typeList := NewConfigService().StructTypeList()
	fmt.Println(typeList)
	c.Render(ctx, "add", gin.H{"root_id": rootId, "root_name": rootName, "typeList": typeList})
}

func (c *StructController) editRender(ctx *gin.Context, model core.IModel) {
	c.Render(ctx, "edit", gin.H{"info": model, "root_id": core.G_STRUCT_ID, "typeList": NewConfigService().StructTypeList()})
}

func (c *StructController) saveBefore(ctx *gin.Context, model core.IModel, operate string) error {
	if operate == "add" || operate == "edit" {
		info := model.(*StructModel)
		if info.ParentId == "" {
			info.ParentId = "0"
		}
		if operate == "add" && info.ParentId == "0" {
			return errors.New("不能添加顶级部门")
		}
		if operate == "edit" && info.Id == core.G_STRUCT_ID && info.ParentId != "0" {
			return errors.New("顶级部门不能修改上级部门")
		}
		if info.ParentId != "0" {
			parentInfo := NewStructModel().New()
			err := core.NewDao(c.IModel).SetField("id,name,parent_name,levels").First(parentInfo, info.ParentId)
			if err != nil {
				return errors.New("上级部门信息不存在")
			}
			info.ParentName = strings.Trim(parentInfo.ParentName+","+parentInfo.Name, ",")
			info.Levels = strings.Trim(parentInfo.Levels+","+parentInfo.Id, ",")
		}
	}
	return nil
}

func (c *StructController) saveAfter(ctx *gin.Context, model core.IModel, operate string, id string) {
	if operate == "add" || operate == "edit" {
		NewStructService().UpdateLevelsInfo(id)
	}
}

func (c *StructController) dropBefore(ctx *gin.Context, model core.IModel) error {
	info, err := core.NewDao(c.IModel).GetInfoByField("parent_id", model.GetId(), "id")
	if err != nil || info.GetId() == "" {
		return nil
	}
	return errors.New("有子组织，无法删除")
}

func (c *StructController) dropAfter(ctx *gin.Context, model core.IModel) {
	_ = NewRoleStructDao().DeleteByStructId(model.GetId())
	_ = NewAdminStructDao().DeleteByStructId(model.GetId())
}

func (c *StructController) Tree(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")
	parent := ctx.DefaultQuery("parent", "0")
	isMult := ctx.DefaultQuery("ismult", "0")
	c.Render(ctx, "tree", gin.H{"struct_id": id, "parent": parent, "ismult": isMult})
}

func (c *StructController) TreeList(ctx *gin.Context) {
	res := NewStructService().TreeList()

	dataScope := services.NewDataScopeFilterByCtx(ctx)
	list := make([]StructModel, 0)
	for _, model := range res {
		if dataScope.CheckByFiled(model.Id, "") {
			list = append(list, model)
		}
	}
	list = NewStructService().ReRootList(list)
	c.Success(ctx, "", list)
}
