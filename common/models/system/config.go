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

type ConfigModel struct {
	Id         string     `json:"id" form:"id"`
	Title      string     `json:"title" form:"title" validate:"required,max=50" label:"标题"`
	Type       string     `json:"type" form:"type" validate:"required,max=50" label:"标识"`
	Style      string     `json:"style" form:"style" validate:"required" label:"类型"`
	Groups     string     `json:"groups" form:"groups" label:"分组"`
	Value      string     `json:"value" form:"value"  label:"配置值"`
	Extra      string     `json:"extra" form:"extra" label:"配置项"`
	Note       string     `json:"note" form:"note" validate:"max=200"  label:"备注"`
	IsSys      string     `db:"is_sys" json:"is_sys" form:"-"`
	CreateTime *time.Time `db:"create_time" json:"create_time" form:"-"`
	UpdateTime *time.Time `db:"update_time" json:"update_time" form:"-"`
}

func (m *ConfigModel) Table() string {
	return "b5net_config"
}

func (m *ConfigModel) GetId() string {
	return m.Id
}

func (m *ConfigModel) DataBase() string {
	return ""
}

// INew 给IModel接口使用创建一个新的结构体
func (m *ConfigModel) INew() core.IModel {
	return m.New()
}

func (m *ConfigModel) New() *ConfigModel {
	return &ConfigModel{}
}

func (m *ConfigModel) NewSlice() *[]ConfigModel {
	return &[]ConfigModel{}
}

var (
	instanceConfigModel *ConfigModel //单例模式
	onceConfigModel     sync.Once
)

// NewConfigModel 单例获取config的结构体
func NewConfigModel() *ConfigModel {
	onceConfigModel.Do(func() {
		instanceConfigModel = &ConfigModel{}
	})
	return instanceConfigModel
}
