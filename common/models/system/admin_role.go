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

type AdminRoleModel struct {
	AdminId string `db:"admin_id" json:"admin_id" form:"admin_id" `
	RoleId  string `db:"role_id" json:"role_id" form:"role_id"`
}

func (m *AdminRoleModel) Table() string {
	return "b5net_admin_role"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *AdminRoleModel) INew() core.IModel {
	return m.New()
}

func (m *AdminRoleModel) GetId() string {
	return ""
}

func (m *AdminRoleModel) DataBase() string {
	return ""
}

func (m *AdminRoleModel) New() *AdminRoleModel {
	return &AdminRoleModel{}
}

func (m *AdminRoleModel) NewSlice() *[]AdminRoleModel {
	return &[]AdminRoleModel{}
}

var (
	instanceAdminRoleModel *AdminRoleModel //单例模式
	onceAdminRoleModel     sync.Once
)

// NewAdminRoleModel 单例获取config的结构体
func NewAdminRoleModel() *AdminRoleModel {
	onceAdminRoleModel.Do(func() {
		instanceAdminRoleModel = &AdminRoleModel{}
	})
	return instanceAdminRoleModel
}
