// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"crypto/md5"
	"errors"
	"fmt"
	"rtsys/admin/lib"
	. "rtsys/common/daos/system"
	. "rtsys/common/models/system"
	. "rtsys/common/services/system"
	"rtsys/utils/core"
	"rtsys/utils/tool"
	"rtsys/utils/trans"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Route 定义该控制器的路由
func (c *AdminController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.POST(c.Dispatch("index", false, c.FindList))
	group.GET(c.Dispatch("add", false, c.Add))
	group.POST(c.Dispatch("add", false, c.AddSave))
	group.GET(c.Dispatch("edit", false, c.Edit))
	group.POST(c.Dispatch("edit", false, c.EditSave))
	group.POST(c.Dispatch("drop", false, c.Drop))
	group.POST(c.Dispatch("drop_all", false, c.DropAll))
	group.POST(c.Dispatch("set_status", false, c.SetStatus))
	group.GET(c.Dispatch("tree", false, c.Tree))
}

type AdminController struct {
	lib.Controller
}

// NewAdminController 创建控制并初始化参数
func NewAdminController() *AdminController {
	c := &AdminController{}
	c.Id = "admin"
	c.Group = "system"
	c.IModel = NewAdminModel()

	c.HookAction.IndexAfter = c.indexAfter
	c.HookAction.EditRender = c.editRender
	c.HookAction.SaveBefore = c.saveBefore
	c.HookAction.SaveAfter = c.saveAfter
	c.HookAction.DropAfter = c.dropAfter
	return c
}

func (c *AdminController) Index(ctx *gin.Context) {
	c.Render(ctx, "index", gin.H{"root_id": core.G_ADMIN_ID, "roleList": NewRoleDao().GetLists()})
}

// FindList 获取列表json
func (c *AdminController) FindList(ctx *gin.Context) {
	where, replace := c.parseIndexSearch(ctx)
	c.GetIndex(ctx, NewAdminModel().NewSlice(), where, replace, nil)
}

// 列表特殊搜索条件处理
type adminSearchExtend struct {
	StructId string `json:"struct_id"`
	RoleId   string `json:"role_id"`
	Contains string `json:"contains"`
}

func (c *AdminController) parseIndexSearch(ctx *gin.Context) (where string, replace []any) {
	searchExtend := &adminSearchExtend{}
	_ = ctx.ShouldBindBodyWith(searchExtend, binding.JSON)
	var userIdList, userIdRoleList, userIdStructList []string
	if searchExtend.RoleId != "" && searchExtend.RoleId != "0" {
		userIdRoleList = NewAdminRoleService().GetAdminIdListByRole(searchExtend.RoleId)
		if userIdRoleList == nil || len(userIdRoleList) < 1 {
			where = "id = ?"
			replace = append(replace, "0")
			return
		}
	}
	if searchExtend.StructId != "" && searchExtend.StructId != "0" {
		structList := make([]string, 0)

		if searchExtend.Contains == "1" {
			if searchExtend.StructId != core.G_STRUCT_ID {
				structList = NewStructService().GetChildAllIdList(searchExtend.StructId)
				structList = append(structList, searchExtend.StructId)
			}
		} else {
			structList = append(structList, searchExtend.StructId)
		}

		if structList != nil && len(structList) > 0 {
			//获取组织下的用户
			userIdStructList = NewAdminStructService().GetAdminIdListByStructList(structList)
			if len(userIdStructList) < 1 {
				where = "id = ?"
				replace = append(replace, "0")
				return
			}
		}
	}
	if len(userIdRoleList) > 0 || len(userIdStructList) > 0 {
		userIdList = tool.UniqueArrStr(tool.MergeArrayStr(userIdRoleList, userIdStructList))
		where = "`id` in (?)"
		replace = append(replace, userIdList)
	}
	return
}

func (c *AdminController) indexAfter(ctx *gin.Context, list []any) []any {
	reList := make([]any, 0)
	for _, item := range list {
		info := item.(AdminModel)
		structName := ""
		posName := ""
		structList := NewAdminStructService().GetStructIdListByAdmin(info.Id)
		roleList := NewAdminRoleService().GetRoleIdListByAdmin(info.Id)
		posList := NewAdminPosService().GetAdminPosIdList(info.Id)
		if len(structList) > 0 {
			structName = NewStructService().GetName(structList[0], true)
		}
		if len(posList) > 0 {
			posName = NewPositionService().GetName(posList[0])
		}
		roleNameList := make([]string, 0)
		for _, roleId := range roleList {
			roleName := NewRoleService().GetName(roleId)
			if roleName != "" {
				roleNameList = append(roleNameList, roleName)
			}
		}
		model := trans.StructToMap(info, true)
		model["struct_name"] = structName
		model["pos_name"] = posName
		model["role_name"] = strings.Join(roleNameList, ",")
		reList = append(reList, model)
	}
	return reList
}

func (c *AdminController) Add(ctx *gin.Context) {
	c.Render(ctx, "add", gin.H{"roleList": NewRoleDao().GetLists(), "posList": NewPositionDao().GetLists()})
}

func (c *AdminController) editRender(ctx *gin.Context, model core.IModel) {
	roleIds := NewAdminRoleService().GetRoleIdListByAdmin(model.GetId())
	structIds := NewAdminStructService().GetStructIdListByAdmin(model.GetId())
	structId := ""
	structName := ""
	if len(structIds) > 0 {
		structId = structIds[0]
		structName = NewStructService().GetName(structId, false)
	}
	posIds := NewAdminPosService().GetAdminPosIdList(model.GetId())
	posId := ""
	if len(posIds) > 0 {
		posId = posIds[0]
	}
	c.Render(ctx, "edit", gin.H{"info": model, "roleList": NewRoleDao().GetLists(), "posList": NewPositionDao().GetLists(), "roleIds": roleIds, "structId": structId, "structName": structName, "posId": posId})
}

type adminExtend struct {
	Roles  string `json:"roles" validate:"required" label:"所属角色"`
	Struct string `json:"struct" validate:"required" label:"组织架构"`
	Pos    string `json:"pos"`
	Pwd    string `json:"pwd" validate:"omitempty,min=6,max=12" label:"密码"`
}

// saveBefore 保存前的字段验证
func (c *AdminController) saveBefore(ctx *gin.Context, model core.IModel, operate string) error {
	if operate != "add" && operate != "edit" {
		return nil
	}
	extend := &adminExtend{}
	err := ctx.ShouldBindBodyWith(extend, binding.JSON)
	if err != nil {
		return errors.New("数据绑定错误")
	}
	err1 := core.G_Validate.Struct(extend)
	if err1 != nil {
		return errors.New(core.G_Validate.GetError(err1))
	}
	if operate == "add" && extend.Pwd == "" {
		return errors.New("请输入密码")
	}
	info := model.(*AdminModel)
	if extend.Pwd != "" {
		info.Password = fmt.Sprintf("%x", md5.Sum([]byte(extend.Pwd)))
	}
	if info.RealName == "" {
		info.RealName = info.UserName
	}
	return nil
}

// saveAfter 保存后更新其他操作
func (c *AdminController) saveAfter(ctx *gin.Context, model core.IModel, operate string, id string) {
	if operate != "add" && operate != "edit" {
		return
	}
	extend := &adminExtend{}
	_ = ctx.ShouldBindBodyWith(extend, binding.JSON)
	_ = NewAdminStructService().UpdateAdminStruct(id, extend.Struct)
	_ = NewAdminRoleService().UpdateAdminRole(id, extend.Roles)
	_ = NewAdminPosService().UpdateAdminPos(id, extend.Pos)
}

// dropAfter 删除信息后，删除管理数据
func (c *AdminController) dropAfter(ctx *gin.Context, model core.IModel) {
	_ = NewAdminStructDao().DeleteByAdminId(model.GetId()) //管理员组织架构
	_ = NewAdminRoleDao().DeleteByAdminId(model.GetId())   //管理员角色
	_ = NewAdminPosDao().DeleteByAdminId(model.GetId())    //管理员岗位
}

func (c *AdminController) Tree(ctx *gin.Context) {
	c.Render(ctx, "tree", gin.H{"user_ids": ctx.DefaultQuery("ids", ""), "mult": ctx.DefaultQuery("mult", "0")})
}
