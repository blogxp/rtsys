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

type MediaModel struct {
	Id         string     `json:"id" form:"id" `
	Img        string     `json:"img" form:"img"`
	Imgs       string     `json:"imgs" form:"imgs"`
	Crop       string     `json:"crop" form:"crop" `
	Video      string     `json:"video" form:"video"`
	File       string     `json:"file" form:"file"`
	Files      string     `json:"files" form:"files"`
	CreateTime *time.Time `db:"create_time" json:"create_time" form:"-"`
	UpdateTime *time.Time `db:"update_time" json:"update_time" form:"-"`
}

func (m *MediaModel) Table() string {
	return "demo_media"
}

func (m *MediaModel) INew() core.IModel {
	return m.New()
}
func (m *MediaModel) GetId() string {
	return m.Id
}

func (m *MediaModel) DataBase() string {
	return ""
}

func (m *MediaModel) New() *MediaModel {
	return &MediaModel{}
}

func (m *MediaModel) NewSlice() *[]MediaModel {
	return &[]MediaModel{}
}

var (
	instanceMediaModel *MediaModel //单例模式
	onceMediaModel     sync.Once
)

// NewMediaModel 单例获取config的结构体
func NewMediaModel() *MediaModel {
	onceMediaModel.Do(func() {
		instanceMediaModel = &MediaModel{}
	})
	return instanceMediaModel
}
