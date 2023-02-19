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

type StructModel struct {
	Id         string     `json:"id" form:"id" `
	Name       string     `json:"name" form:"name" validate:"required" label:"组织名称"`
	Type       string     `json:"type" form:"type" validate:"required" label:"组织类型"`
	ParentId   string     `db:"parent_id" json:"parent_id" form:"parent_id" validate:"required" label:"所属父级"`
	Levels     string     `json:"levels" form:"-"`
	ParentName string     `db:"parent_name" json:"parent_name" form:"parent_name" form:"-"`
	ListSort   string     `db:"list_sort" json:"list_sort" form:"list_sort" validate:"required" label:"显示顺序"`
	Leader     string     `json:"leader" form:"leader"  validate:"max=20" label:"负责人"`
	Phone      string     `json:"phone" form:"phone" validate:"max=20" label:"联系电话"`
	Note       string     `json:"note" form:"note" validate:"max=200" label:"备注"`
	Status     string     `json:"status" form:"status" validate:"required,oneof=0 1" label:"状态"`
	CreateTime *time.Time `db:"create_time" json:"create_time" form:"-"`
	UpdateTime *time.Time `db:"update_time" json:"update_time" form:"-"`
}

func (m *StructModel) Table() string {
	return "b5net_struct"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *StructModel) INew() core.IModel {
	return m.New()
}

func (m *StructModel) GetId() string {
	return m.Id
}

func (m *StructModel) DataBase() string {
	return ""
}

func (m *StructModel) New() *StructModel {
	return &StructModel{}
}

func (m *StructModel) NewSlice() *[]StructModel {
	return &[]StructModel{}
}

var (
	instanceStructModel *StructModel //单例模式
	onceStructModel     sync.Once
)

// NewStructModel 单例获取config的结构体
func NewStructModel() *StructModel {
	onceStructModel.Do(func() {
		instanceStructModel = &StructModel{}
	})
	return instanceStructModel
}
