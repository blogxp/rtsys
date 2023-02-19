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

type LoginLogModel struct {
	Id            string     `json:"id" form:"id" `
	LoginName     string     `db:"login_name" json:"login_name" form:"login_name" label:"登陆账号"`
	IpAddr        string     `db:"ip_addr" json:"ip_addr" form:"ip_addr" label:"IP地址"`
	LoginLocation string     `db:"login_location" json:"login_location" form:"login_location" label:"登录地点"`
	Browser       string     `json:"browser" form:"browser" label:"浏览器"`
	Os            string     `json:"os" form:"os" label:"操作系统"`
	Net           string     `json:"net" form:"net" label:"网络"`
	Msg           string     `json:"msg" form:"msg" label:"提示消息"`
	Status        string     `json:"status" form:"status" label:"状态"`
	CreateTime    *time.Time `db:"create_time" json:"create_time" form:"-"`
	UpdateTime    *time.Time `db:"update_time" json:"update_time" form:"-"`
}

func (m *LoginLogModel) Table() string {
	return "b5net_login_log"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *LoginLogModel) INew() core.IModel {
	return m.New()
}

func (m *LoginLogModel) GetId() string {
	return m.Id
}

func (m *LoginLogModel) DataBase() string {
	return ""
}

func (m *LoginLogModel) New() *LoginLogModel {
	return &LoginLogModel{}
}

func (m *LoginLogModel) NewSlice() *[]LoginLogModel {
	return &[]LoginLogModel{}
}

var (
	instanceLoginLogModel *LoginLogModel //单例模式
	onceLoginLogModel     sync.Once
)

// NewLoginLogModel 单例获取config的结构体
func NewLoginLogModel() *LoginLogModel {
	onceLoginLogModel.Do(func() {
		instanceLoginLogModel = &LoginLogModel{}
	})
	return instanceLoginLogModel
}
