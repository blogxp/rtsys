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

type PositionModel struct {
	Id         string     `json:"id" form:"id" `
	Name       string     `json:"name" form:"name" validate:"required" label:"岗位名称"`
	PosKey     string     `db:"pos_key" json:"pos_key" form:"pos_key" validate:"required" label:"岗位标识"`
	ListSort   string     `db:"list_sort" json:"list_sort" form:"list_sort" validate:"required" label:"显示顺序"`
	Note       string     `json:"note" form:"note" validate:"max=200" label:"备注"`
	Status     string     `json:"status" form:"status" validate:"required,oneof=0 1" label:"状态"`
	CreateTime *time.Time `db:"create_time" json:"create_time" form:"-"`
	UpdateTime *time.Time `db:"update_time" json:"update_time" form:"-"`
}

func (m *PositionModel) Table() string {
	return "b5net_position"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *PositionModel) INew() core.IModel {
	return m.New()
}

func (m *PositionModel) GetId() string {
	return m.Id
}

func (m *PositionModel) DataBase() string {
	return ""
}

func (m *PositionModel) New() *PositionModel {
	return &PositionModel{}
}

func (m *PositionModel) NewSlice() *[]PositionModel {
	return &[]PositionModel{}
}

var (
	instancePositionModel *PositionModel //单例模式
	oncePositionModel     sync.Once
)

// NewPositionModel 单例获取config的结构体
func NewPositionModel() *PositionModel {
	oncePositionModel.Do(func() {
		instancePositionModel = &PositionModel{}
	})
	return instancePositionModel
}
