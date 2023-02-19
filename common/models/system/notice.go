// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"database/sql"
	"rtsys/utils/core"
	"sync"
	"time"
)

type NoticeModel struct {
	Id         string         `json:"id" form:"id" `
	Title      string         `json:"title" form:"title" validate:"required"`
	Type       string         `json:"type" form:"type" validate:"required"`
	Desc       sql.NullString `json:"desc" form:"desc" `
	Content    string         `json:"content" form:"content"`
	Status     string         `json:"status" form:"status"`
	CreateTime *time.Time     `db:"create_time" json:"create_time" form:"-"`
	UpdateTime *time.Time     `db:"update_time" json:"update_time" form:"-"`
}

func (m *NoticeModel) Table() string {
	return "b5net_notice"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *NoticeModel) INew() core.IModel {
	return m.New()
}
func (m *NoticeModel) GetId() string {
	return m.Id
}

func (m *NoticeModel) DataBase() string {
	return ""
}

func (m *NoticeModel) New() *NoticeModel {
	return &NoticeModel{}
}

func (m *NoticeModel) NewSlice() *[]NoticeModel {
	return &[]NoticeModel{}
}

var (
	instanceNoticeModel *NoticeModel //单例模式
	onceNoticeModel     sync.Once
)

// NewNoticeModel 单例获取config的结构体
func NewNoticeModel() *NoticeModel {
	onceNoticeModel.Do(func() {
		instanceNoticeModel = &NoticeModel{}
	})
	return instanceNoticeModel
}
