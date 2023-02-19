// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	. "rtsys/common/daos/system"
	. "rtsys/common/models/system"
	"sync"
)

type AdminService struct {
	Dao *AdminDao
}

var (
	instanceAdminService *AdminService //单例的对象
	onceAdminService     sync.Once
)

func NewAdminService() *AdminService {
	onceAdminService.Do(func() {
		instanceAdminService = &AdminService{Dao: NewAdminDao()}
	})
	return instanceAdminService
}

func (s *AdminService) GetName(id string) string {
	if id == "0" || id == "" {
		return ""
	}
	info := s.GetInfo(id)
	if info == nil {
		return ""
	}
	return info.RealName
}

func (s *AdminService) GetInfo(id string) *AdminModel {
	return s.Dao.GetInfoById(id)
}
