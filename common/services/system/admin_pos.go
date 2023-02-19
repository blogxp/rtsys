// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"errors"
	. "rtsys/common/daos/system"
	"strings"
	"sync"
)

type AdminPosService struct {
	Dao *AdminPosDao
}

var (
	instanceAdminPosService *AdminPosService //单例的对象
	onceAdminPosService     sync.Once
)

func NewAdminPosService() *AdminPosService {
	onceAdminPosService.Do(func() {
		instanceAdminPosService = &AdminPosService{Dao: NewAdminPosDao()}
	})
	return instanceAdminPosService
}

// GetAdminPosIdList 获取某个管理员的岗位ID数组
func (s *AdminPosService) GetAdminPosIdList(adminId string) (reList []string) {
	list, err := s.Dao.GetAdminPosList(adminId)
	if err != nil {
		return
	}
	for _, v := range *list {
		reList = append(reList, v.PosId)
	}
	return
}

// UpdateAdminPos 更新管理员的岗位节点
func (s *AdminPosService) UpdateAdminPos(adminId string, posIds string) error {
	if adminId == "" {
		return nil
	}
	err := s.Dao.DeleteByAdminId(adminId)
	if err != nil {
		return errors.New("删除原节点失败")
	}
	if posIds == "" {
		return nil
	}
	s.Dao.InsertList(adminId, strings.Split(posIds, ","))
	return nil
}
