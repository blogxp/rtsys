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
	"rtsys/utils/core"
	"rtsys/utils/export"
	"rtsys/utils/trans"
	"rtsys/utils/types"

	"github.com/gin-gonic/gin"
)

// Route 定义该控制器的路由
func (c *PositionController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.POST(c.Dispatch("index", false, c.FindList))
	group.GET(c.Dispatch("add", false, c.Add))
	group.POST(c.Dispatch("add", false, c.AddSave))
	group.GET(c.Dispatch("edit", false, c.Edit))
	group.POST(c.Dispatch("edit", false, c.EditSave))
	group.POST(c.Dispatch("drop", false, c.Drop))
	group.POST(c.Dispatch("drop_all", false, c.DropAll))
	group.POST(c.Dispatch("set_status", false, c.SetStatus))
}

type PositionController struct {
	lib.Controller
}

// NewPositionController 创建控制并初始化参数
func NewPositionController() *PositionController {
	c := &PositionController{}
	c.Id = "position"
	c.Group = "system"
	c.IModel = NewPositionModel()

	c.HookAction.SaveBefore = c.saveBefore
	c.HookAction.DropAfter = c.dropAfter
	c.HookAction.ExportBefore = c.exportBefore
	return c
}

// FindList 获取列表json
func (c *PositionController) FindList(ctx *gin.Context) {
	c.GetIndex(ctx, NewPositionModel().NewSlice(), "", nil, nil)
}

// saveBefore 保存前验证编码唯一性
func (c *PositionController) saveBefore(ctx *gin.Context, model core.IModel, operate string) error {
	if operate == "add" || operate == "edit" {
		model := model.(*PositionModel)
		info, err := core.NewDao(c.IModel).GetInfoByField("pos_key", model.PosKey, "")
		if err != nil {
			return nil
		}
		if (operate == "edit" && info.GetId() != model.GetId()) || operate == "add" {
			return errors.New("岗位编码已存在")
		}
	}
	return nil
}

func (c *PositionController) dropAfter(ctx *gin.Context, model core.IModel) {
	_ = NewAdminPosDao().DeleteByPosId(model.GetId())
}

func (c *PositionController) exportBefore(ctx *gin.Context, list []any) *export.ExcelExport {
	excel := &export.ExcelExport{}
	excel.Labels = []types.KeyVal{
		{Key: "name", Value: "名称"},
		{Key: "pos_key", Value: "标识"},
		{Key: "status", Value: "状态"},
		{Key: "create_time", Value: "创建时间"},
		{Key: "update_time", Value: "更新时间"},
	}
	excel.List = trans.StructListToMapList(list, true)
	return excel
}
