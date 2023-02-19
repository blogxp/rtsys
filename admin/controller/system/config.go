// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"errors"
	"rtsys/admin/lib"
	. "rtsys/common/models/system"
	. "rtsys/common/services/system"
	"rtsys/utils/core"
	"rtsys/utils/export"
	"rtsys/utils/trans"
	"rtsys/utils/types"

	"github.com/gin-gonic/gin"
)

// Route 定义该控制器的路由
func (c *ConfigController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.POST(c.Dispatch("index", false, c.FindList))
	group.GET(c.Dispatch("add", false, c.Add))
	group.POST(c.Dispatch("add", false, c.AddSave))
	group.GET(c.Dispatch("edit", false, c.Edit))
	group.POST(c.Dispatch("edit", false, c.EditSave))
	group.POST(c.Dispatch("drop", false, c.Drop))
	group.POST(c.Dispatch("drop_all", false, c.DropAll))
	group.GET(c.Dispatch("site", false, c.Site))
	group.POST(c.Dispatch("site", false, c.SiteSave))
}

type ConfigController struct {
	lib.Controller
}

// NewConfigController 创建控制并初始化参数
func NewConfigController() *ConfigController {
	c := &ConfigController{}
	c.Id = "config"
	c.Group = "system"
	c.IModel = NewConfigModel()
	c.HookAction.IndexRender = c.indexRender
	c.HookAction.IndexAfter = c.indexAfter
	c.HookAction.EditRender = c.eddRender
	c.HookAction.SaveBefore = c.saveBefore
	c.HookAction.ExportBefore = c.exportBefore
	return c
}

func (c *ConfigController) Site(ctx *gin.Context) {
	groups := NewConfigService().GetListByGroup()
	c.Render(ctx, "site", gin.H{"groups": groups})
}
func (c *ConfigController) SiteSave(ctx *gin.Context) {
	autokey := &types.AutoKey{}
	err := ctx.ShouldBind(autokey)
	if err != nil {
		c.Error(ctx, "数据提交失败")
		return
	}
	NewConfigService().UpdateByAutoKey(autokey)
	c.Success(ctx, "更新成功")
}

func (c *ConfigController) indexRender(ctx *gin.Context) (map[string]any, error) {
	return gin.H{"styleList": NewConfigService().StyleList(), "groupList": NewConfigService().GroupList()}, nil
}

func (c *ConfigController) indexAfter(ctx *gin.Context, list []any) []any {
	if len(list) > 0 {
		service := NewConfigService()
		groupList := service.GroupList()
		for i := 0; i < len(list); i++ {
			elem := list[i]
			model := elem.(ConfigModel)
			maps := trans.StructToMap(elem, true)
			maps["group_name"] = service.GetGroupName(model.Groups, groupList)
			maps["style_name"] = service.GetStyleName(model.Style)
			list[i] = maps
		}
	}
	return list
}

// FindList 获取列表json
func (c *ConfigController) FindList(ctx *gin.Context) {
	c.GetIndex(ctx, NewConfigModel().NewSlice(), "", nil, []types.KeyVal{{Key: "is_sys", Value: "desc"}})
}

// Add 覆盖
func (c *ConfigController) Add(ctx *gin.Context) {
	c.Render(ctx, "add", gin.H{"styleList": NewConfigService().StyleList(), "groupList": NewConfigService().GroupList()})
}

// eddRender 渲染前操作
func (c *ConfigController) eddRender(ctx *gin.Context, model core.IModel) {
	c.Render(ctx, "edit", gin.H{"styleList": NewConfigService().StyleList(), "groupList": NewConfigService().GroupList(), "info": model})
}

func (c *ConfigController) saveBefore(ctx *gin.Context, model core.IModel, operate string) error {
	if operate == "add" || operate == "edit" {
		model := model.(*ConfigModel)
		if operate == "add" {
			model.IsSys = "0"
		}
		info, err := core.NewDao(c.IModel).GetInfoByField("type", model.Type, "")
		if err != nil {
			return nil
		}
		if (operate == "edit" && info.GetId() != model.GetId()) || operate == "add" {
			return errors.New("标识重复了")
		}
	}
	return nil
}

func (c *ConfigController) exportBefore(ctx *gin.Context, list []any) *export.ExcelExport {
	excel := &export.ExcelExport{}
	excel.Labels = []types.KeyVal{
		{Key: "title", Value: "配置标题"},
		{Key: "type", Value: "配置标识"},
		{Key: "style_name", Value: "配置类型"},
		{Key: "value", Value: "配置值"},
		{Key: "extra", Value: "配置项"},
		{Key: "update_time", Value: "更新时间"},
	}
	excel.List = trans.StructListToMapList(list, true)
	return excel
}
