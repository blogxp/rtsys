// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package __GROUP__

import (
	"rtsys/utils/core"
	"sync"
	__TIME__
)

type __MODEL__ struct {
__FIELD__
}

func (m *__MODEL__) Table() string {
	return "__TABLE__"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *__MODEL__) INew() core.IModel {
	return m.New()
}

func (m *__MODEL__) GetId() string {
	return __PRIMARY__
}

func (m *__MODEL__) DataBase() string {
	return ""
}

func (m *__MODEL__) New() *__MODEL__ {
	return &__MODEL__{}
}

func (m *__MODEL__) NewSlice() *[]__MODEL__ {
	return &[]__MODEL__{}
}

var (
	instance__MODEL__ *__MODEL__ //单例模式
	once__MODEL__     sync.Once
)

// New__MODEL__ 单例获取
func New__MODEL__() *__MODEL__ {
	once__MODEL__.Do(func() {
		instance__MODEL__ = &__MODEL__{}
	})
	return instance__MODEL__
}
