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

type AdminModel struct {
	Id         string     `json:"id" form:"id"`
	UserName   string     `json:"username" form:"username" validate:"required,max=30" label:"登录名称"`
	Password   string     `json:"password" form:"-"`
	RealName   string     `json:"realname" form:"realname"`
	Status     string     `db:"status" json:"status" form:"-" validate:"required,oneof=0 1" label:"状态"`
	Note       string     `json:"note" form:"note"`
	CreateTime *time.Time `db:"create_time" json:"create_time" form:"-"`
	UpdateTime *time.Time `db:"update_time" json:"update_time" form:"-"`
}

func (m *AdminModel) Table() string {
	return "b5net_admin"
}

func (m *AdminModel) GetId() string {
	return m.Id
}

func (m *AdminModel) DataBase() string {
	return ""
}

// INew 给IModel接口使用创建一个新的结构体
func (m *AdminModel) INew() core.IModel {
	return m.New()
}

func (m *AdminModel) New() *AdminModel {
	return &AdminModel{}
}

func (m *AdminModel) NewSlice() *[]AdminModel {
	return &[]AdminModel{}
}

var (
	instanceAdminModel *AdminModel //单例模式
	onceAdminModel     sync.Once
)

// NewAdminModel 单例获取config的结构体
func NewAdminModel() *AdminModel {
	onceAdminModel.Do(func() {
		instanceAdminModel = &AdminModel{}
	})
	return instanceAdminModel
}
