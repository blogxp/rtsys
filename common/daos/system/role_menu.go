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

type RoleMenuDao struct {
	Model *RoleMenuModel
}

var (
	instanceRoleMenuDao *RoleMenuDao //单例的对象
	onceRoleMenuDao     sync.Once
)

func NewRoleMenuDao() *RoleMenuDao {
	onceRoleMenuDao.Do(func() {
		instanceRoleMenuDao = &RoleMenuDao{Model: NewRoleMenuModel()}
	})
	return instanceRoleMenuDao
}

// GetRoleMenuList 获取某个角色的菜单列表
func (d *RoleMenuDao) GetRoleMenuList(roleId string) (*[]RoleMenuModel, error) {
	if roleId == "" {
		return nil, errors.New("参数丢失")
	}
	lists := d.Model.NewSlice()
	err := core.NewDao(d.Model).Lists(lists, "`role_id` = ?", roleId)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

// DeleteByRoleId 删除某个角色的菜单权限节点
func (d *RoleMenuDao) DeleteByRoleId(roleId string) error {
	if roleId == "" {
		return nil
	}
	_, err := core.NewDao(d.Model).Delete("`role_id` = ?", roleId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByMenuId 删除某个菜单的角色权限节点
func (d *RoleMenuDao) DeleteByMenuId(menuId string) error {
	if menuId == "" {
		return nil
	}
	_, err := core.NewDao(d.Model).Delete("`menu_id` = ?", menuId)
	if err != nil {
		return err
	}
	return nil
}

// InsertList 插入角色菜单节点
func (d *RoleMenuDao) InsertList(roleId string, menuIds []string) int64 {
	var affected int64 = 0
	for _, id := range menuIds {
		if id == "" {
			continue
		}
		err := core.NewDao(d.Model).InsertNoId([]string{"role_id", "menu_id"}, []any{roleId, id})
		if err == nil {
			affected++
		}
	}
	return affected
}
