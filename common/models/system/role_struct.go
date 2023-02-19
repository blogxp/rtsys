// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"rtsys/utils/core"
	"sync"
)

type RoleStructModel struct {
	RoleId   string `db:"role_id" json:"role_id" form:"role_id" `
	StructId string `db:"struct_id" json:"struct_id" form:"struct_id"`
}

func (m *RoleStructModel) Table() string {
	return "b5net_role_struct"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *RoleStructModel) INew() core.IModel {
	return m.New()
}

func (m *RoleStructModel) GetId() string {
	return ""
}

func (m *RoleStructModel) DataBase() string {
	return ""
}

func (m *RoleStructModel) New() *RoleStructModel {
	return &RoleStructModel{}
}

func (m *RoleStructModel) NewSlice() *[]RoleStructModel {
	return &[]RoleStructModel{}
}

var (
	instanceRoleStructModel *RoleStructModel //单例模式
	onceRoleStructModel     sync.Once
)

// NewRoleStructModel 单例获取config的结构体
func NewRoleStructModel() *RoleStructModel {
	onceRoleStructModel.Do(func() {
		instanceRoleStructModel = &RoleStructModel{}
	})
	return instanceRoleStructModel
}
