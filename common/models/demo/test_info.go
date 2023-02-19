// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package demo

import (
	"rtsys/utils/core"
	"sync"
	"time"
)

type TestInfoModel struct {
	Id           string     `json:"id" form:"id"`                                           //
	StructId     string     `db:"struct_id" json:"struct_id" form:"-"`                      // 组织ID
	StructLevels string     `db:"struct_levels" json:"struct_levels" form:"-"`              // 组织ID的levels
	UserId       string     `db:"user_id" json:"user_id" form:"-"`                          // 用户ID
	Name         string     `json:"name" form:"name" validate:"required,max=30" label:"标题"` // 标题
	Status       string     `json:"status" form:"status" validate:"required" label:"状态"`    // 状态
	Remark       string     `json:"remark" form:"remark"`                                   // 介绍
	CreateTime   *time.Time `db:"create_time" json:"create_time" form:"-"`                  //
	UpdateTime   *time.Time `db:"update_time" json:"update_time" form:"-"`                  //

}

func (m *TestInfoModel) Table() string {
	return "test_info"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *TestInfoModel) INew() core.IModel {
	return m.New()
}

func (m *TestInfoModel) GetId() string {
	return m.Id
}

func (m *TestInfoModel) DataBase() string {
	return ""
}

func (m *TestInfoModel) New() *TestInfoModel {
	return &TestInfoModel{}
}

func (m *TestInfoModel) NewSlice() *[]TestInfoModel {
	return &[]TestInfoModel{}
}

var (
	instanceTestInfoModel *TestInfoModel //单例模式
	onceTestInfoModel     sync.Once
)

// NewTestInfoModel 单例获取
func NewTestInfoModel() *TestInfoModel {
	onceTestInfoModel.Do(func() {
		instanceTestInfoModel = &TestInfoModel{}
	})
	return instanceTestInfoModel
}
