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

type AdminRoleService struct {
	Dao *AdminRoleDao
}

var (
	instanceAdminRoleService *AdminRoleService //单例的对象
	onceAdminRoleService     sync.Once
)

func NewAdminRoleService() *AdminRoleService {
	onceAdminRoleService.Do(func() {
		instanceAdminRoleService = &AdminRoleService{Dao: NewAdminRoleDao()}
	})
	return instanceAdminRoleService
}

// GetRoleIdListByAdmin 获取某个管理员的角色ID数组
func (s *AdminRoleService) GetRoleIdListByAdmin(adminId string) (reList []string) {
	list, err := s.Dao.GetRoleListByAdmin(adminId)
	if err != nil {
		return
	}
	for _, v := range *list {
		reList = append(reList, v.RoleId)
	}
	return
}

// GetAdminIdListByRole 获取某个角色的管理员ID数组
func (s *AdminRoleService) GetAdminIdListByRole(roleId string) (reList []string) {
	list, err := s.Dao.GetAdminListByRole(roleId)
	if err != nil {
		return
	}
	for _, v := range *list {
		reList = append(reList, v.AdminId)
	}
	return
}

// UpdateAdminRole 更新管理员的角色节点
func (s *AdminRoleService) UpdateAdminRole(adminId string, roleIds string) error {
	if adminId == "" {
		return nil
	}
	err := s.Dao.DeleteByAdminId(adminId)
	if err != nil {
		return errors.New("删除原节点失败")
	}
	if roleIds == "" {
		return nil
	}
	s.Dao.InsertList(adminId, strings.Split(roleIds, ","))
	return nil
}
