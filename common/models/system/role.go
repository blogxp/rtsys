// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"rtsys/utils/core"
	"sync"
	"time"
)

type RoleModel struct {
	Id         string     `json:"id" form:"id" `
	Name       string     `json:"name" form:"name" validate:"required" label:"岗位名称"`
	RoleKey    string     `db:"role_key" json:"role_key" form:"role_key" validate:"required" label:"角色标识"`
	DataScope  string     `db:"data_scope" json:"data_scope" form:"data_scope" label:"数据范围"`
	ListSort   string     `db:"list_sort" json:"list_sort" form:"list_sort" validate:"required" label:"显示顺序"`
	Note       string     `json:"note" form:"note" validate:"max=200" label:"备注"`
	Status     string     `json:"status" form:"status" validate:"required,oneof=0 1" label:"状态"`
	CreateTime *time.Time `db:"create_time" json:"create_time" form:"-"`
	UpdateTime *time.Time `db:"update_time" json:"update_time" form:"-"`
}

func (m *RoleModel) Table() string {
	return "b5net_role"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *RoleModel) INew() core.IModel {
	return m.New()
}

func (m *RoleModel) GetId() string {
	return m.Id
}

func (m *RoleModel) DataBase() string {
	return ""
}

func (m *RoleModel) New() *RoleModel {
	return &RoleModel{}
}

func (m *RoleModel) NewSlice() *[]RoleModel {
	return &[]RoleModel{}
}

var (
	instanceRoleModel *RoleModel //单例模式
	onceRoleModel     sync.Once
)

// NewRoleModel 单例获取config的结构体
func NewRoleModel() *RoleModel {
	onceRoleModel.Do(func() {
		instanceRoleModel = &RoleModel{}
	})
	return instanceRoleModel
}
