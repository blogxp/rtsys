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

type RoleStructDao struct {
	Model *RoleStructModel
}

var (
	instanceRoleStructDao *RoleStructDao //单例的对象
	onceRoleStructDao     sync.Once
)

func NewRoleStructDao() *RoleStructDao {
	onceRoleStructDao.Do(func() {
		instanceRoleStructDao = &RoleStructDao{Model: NewRoleStructModel()}
	})
	return instanceRoleStructDao
}

// GetRoleStructList 获取某个角色的组织列表
func (d *RoleStructDao) GetRoleStructList(roleId string) (*[]RoleStructModel, error) {
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

// DeleteByRoleId 删除某个角色的组织权限节点
func (d *RoleStructDao) DeleteByRoleId(roleId string) error {
	if roleId == "" {
		return nil
	}
	_, err := core.NewDao(d.Model).Delete("`role_id` = ?", roleId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByStructId 删除组织对应的角色信息
func (d *RoleStructDao) DeleteByStructId(structId string) error {
	if structId == "" {
		return nil
	}
	_, err := core.NewDao(d.Model).Delete("`struct_id` = ?", structId)
	if err != nil {
		return err
	}
	return nil
}

// InsertList 插入角色组织节点
func (d *RoleStructDao) InsertList(roleId string, structIds []string) int64 {
	var affected int64 = 0
	for _, id := range structIds {
		if id == "" {
			continue
		}
		err := core.NewDao(d.Model).InsertNoId([]string{"role_id", "struct_id"}, []any{roleId, id})
		if err == nil {
			affected++
		}
	}
	return affected
}
