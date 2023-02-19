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

type MenuModel struct {
	Id         string     `json:"id" form:"id"`
	Name       string     `json:"name" form:"name"  validate:"required,max=30" label:"菜单名称"`
	ParentId   string     `db:"parent_id" json:"parent_id" form:"parent_id" label:"父菜单ID"`
	ListSort   string     `db:"list_sort" json:"list_sort" form:"list_sort"  validate:"required" label:"显示顺序"`
	Url        string     `json:"url" form:"url" label:"请求地址"`
	Target     string     `json:"target" form:"target"  validate:"required,oneof=0 1" label:"打开方式"`
	Type       string     `json:"type" form:"type"  validate:"required" label:"菜单类型"`
	Status     string     `json:"status"  form:"status"  validate:"required,oneof=0 1" label:"菜单状态"`
	IsRefresh  string     `db:"is_refresh" json:"is_refresh"  validate:"required,oneof=0 1" label:"是否刷新"`
	Perms      string     `json:"perms" form:"perms" label:"权限标识"`
	Icon       string     `json:"icon" form:"icon" label:"菜单图标"`
	Note       string     `json:"note" form:"note"`
	CreateTime *time.Time `db:"create_time" json:"create_time" form:"-"`
	UpdateTime *time.Time `db:"update_time" json:"update_time" form:"-"`
}

func (m *MenuModel) Table() string {
	return "b5net_menu"
}

func (m *MenuModel) GetId() string {
	return m.Id
}

func (m *MenuModel) DataBase() string {
	return ""
}

// INew 给IModel接口使用创建一个新的结构体
func (m *MenuModel) INew() core.IModel {
	return m.New()
}

func (m *MenuModel) New() *MenuModel {
	return &MenuModel{}
}

func (m *MenuModel) NewSlice() *[]MenuModel {
	return &[]MenuModel{}
}

var (
	instanceMenuModel *MenuModel //单例模式
	onceMenuModel     sync.Once
)

// NewMenuModel 单例获取config的结构体
func NewMenuModel() *MenuModel {
	onceMenuModel.Do(func() {
		instanceMenuModel = &MenuModel{}
	})
	return instanceMenuModel
}
