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

type AppTokenModel struct {
	Token   string     `json:"token"`
	Type    string     `json:"type"`
	Plat    string     `json:"plat"`
	UserId  string     `db:"user_id" json:"user_id"`
	Extend  string     `json:"extend"`
	ExpTime *time.Time `db:"exp_time" json:"exp_time" form:"-"`
}

func (m *AppTokenModel) Table() string {
	return "b5net_app_token"
}

func (m *AppTokenModel) GetId() string {
	return m.Token
}

func (m *AppTokenModel) DataBase() string {
	return ""
}

// INew 给IModel接口使用创建一个新的结构体
func (m *AppTokenModel) INew() core.IModel {
	return m.New()
}

func (m *AppTokenModel) New() *AppTokenModel {
	return &AppTokenModel{}
}

var (
	instanceAppTokenModel *AppTokenModel //单例模式
	onceAppTokenModel     sync.Once
)

// NewAppTokenModel 单例获取AppToken的结构体
func NewAppTokenModel() *AppTokenModel {
	onceAppTokenModel.Do(func() {
		instanceAppTokenModel = &AppTokenModel{}
	})
	return instanceAppTokenModel
}
