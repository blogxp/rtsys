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

type AdminPosModel struct {
	AdminId string `db:"admin_id" json:"admin_id" form:"admin_id" `
	PosId   string `db:"pos_id" json:"pos_id" form:"pos_id"`
}

func (m *AdminPosModel) Table() string {
	return "b5net_admin_pos"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *AdminPosModel) INew() core.IModel {
	return m.New()
}

func (m *AdminPosModel) GetId() string {
	return ""
}

func (m *AdminPosModel) DataBase() string {
	return ""
}

func (m *AdminPosModel) New() *AdminPosModel {
	return &AdminPosModel{}
}

func (m *AdminPosModel) NewSlice() *[]AdminPosModel {
	return &[]AdminPosModel{}
}

var (
	instanceAdminPosModel *AdminPosModel //单例模式
	onceAdminPosModel     sync.Once
)

// NewAdminPosModel 单例获取config的结构体
func NewAdminPosModel() *AdminPosModel {
	onceAdminPosModel.Do(func() {
		instanceAdminPosModel = &AdminPosModel{}
	})
	return instanceAdminPosModel
}
