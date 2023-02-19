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

type AdminStructModel struct {
	AdminId  string `db:"admin_id" json:"admin_id" form:"admin_id" `
	StructId string `db:"struct_id" json:"struct_id" form:"struct_id"`
}

func (m *AdminStructModel) Table() string {
	return "b5net_admin_struct"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *AdminStructModel) INew() core.IModel {
	return m.New()
}

func (m *AdminStructModel) GetId() string {
	return ""
}

func (m *AdminStructModel) DataBase() string {
	return ""
}

func (m *AdminStructModel) New() *AdminStructModel {
	return &AdminStructModel{}
}

func (m *AdminStructModel) NewSlice() *[]AdminStructModel {
	return &[]AdminStructModel{}
}

var (
	instanceAdminStructModel *AdminStructModel //单例模式
	onceAdminStructModel     sync.Once
)

// NewAdminStructModel 单例获取config的结构体
func NewAdminStructModel() *AdminStructModel {
	onceAdminStructModel.Do(func() {
		instanceAdminStructModel = &AdminStructModel{}
	})
	return instanceAdminStructModel
}
