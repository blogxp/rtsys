// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	. "rtsys/common/models/system"
	"sync"
)

type NoticeDao struct {
	Model *NoticeModel
}

var (
	instanceNoticeDao *NoticeDao //单例的对象
	onceNoticeDao     sync.Once
)

func NewNoticeDao() *NoticeDao {
	onceNoticeDao.Do(func() {
		instanceNoticeDao = &NoticeDao{Model: NewNoticeModel()}
	})
	return instanceNoticeDao
}
