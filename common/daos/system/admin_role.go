// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"errors"
	. "rtsys/common/models/system"
	"rtsys/utils/core"
	"sync"
)

type AdminRoleDao struct {
	Model *AdminRoleModel
}

var (
	instanceAdminRoleDao *AdminRoleDao //单例的对象
	onceAdminRoleDao     sync.Once
)

func NewAdminRoleDao() *AdminRoleDao {
	onceAdminRoleDao.Do(func() {
		instanceAdminRoleDao = &AdminRoleDao{Model: NewAdminRoleModel()}
	})
	return instanceAdminRoleDao
}

// GetRoleListByAdmin 获取某个管理员的角色列表
func (d *AdminRoleDao) GetRoleListByAdmin(adminId string) (*[]AdminRoleModel, error) {
	if adminId == "" {
		return nil, errors.New("参数丢失")
	}
	lists := d.Model.NewSlice()
	err := core.NewDao(d.Model).Lists(lists, "`admin_id` = ?", adminId)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

// GetAdminListByRole 获取某个管理员的角色列表
func (d *AdminRoleDao) GetAdminListByRole(adminId string) (*[]AdminRoleModel, error) {
	if adminId == "" {
		return nil, errors.New("参数丢失")
	}
	lists := d.Model.NewSlice()
	err := core.NewDao(d.Model).Lists(lists, "`role_id` = ?", adminId)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

// DeleteByAdminId 删除某个管理员的角色权限节点
func (d *AdminRoleDao) DeleteByAdminId(adminId string) error {
	if adminId == "" {
		return nil
	}
	_, err := core.NewDao(d.Model).Delete("`admin_id` = ?", adminId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByRoleId 删除某个角色的管理员
func (d *AdminRoleDao) DeleteByRoleId(roleId string) error {
	if roleId == "" {
		return nil
	}
	_, err := core.NewDao(d.Model).Delete("`role_id` = ?", roleId)
	if err != nil {
		return err
	}
	return nil
}

// InsertList 插入管理员角色节点
func (d *AdminRoleDao) InsertList(adminId string, roleIds []string) int64 {
	var affected int64 = 0
	for _, id := range roleIds {
		if id == "" {
			continue
		}
		err := core.NewDao(d.Model).InsertNoId([]string{"admin_id", "role_id"}, []any{adminId, id})
		if err == nil {
			affected++
		}
	}
	return affected
}
