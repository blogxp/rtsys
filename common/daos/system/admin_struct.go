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

type AdminStructDao struct {
	Model *AdminStructModel
}

var (
	instanceAdminStructDao *AdminStructDao //单例的对象
	onceAdminStructDao     sync.Once
)

func NewAdminStructDao() *AdminStructDao {
	onceAdminStructDao.Do(func() {
		instanceAdminStructDao = &AdminStructDao{Model: NewAdminStructModel()}
	})
	return instanceAdminStructDao
}

// GetStructListByAdmin 获取某个管理员的组织列表
func (d *AdminStructDao) GetStructListByAdmin(adminId string) (*[]AdminStructModel, error) {
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

// GetAdminListByStruct 获取某个管理员的组织列表
func (d *AdminStructDao) GetAdminListByStruct(structId string) (*[]AdminStructModel, error) {
	if structId == "" {
		return nil, errors.New("参数丢失")
	}
	lists := d.Model.NewSlice()
	err := core.NewDao(d.Model).Lists(lists, "`struct_id` = ?", structId)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

// GetAdminListByStructList 获取某个管理员的组织列表
func (d *AdminStructDao) GetAdminListByStructList(structIds []string) (*[]AdminStructModel, error) {
	if len(structIds) < 1 {
		return nil, errors.New("参数丢失")
	}
	lists := d.Model.NewSlice()
	err := core.NewDao(d.Model).Lists(lists, "`struct_id` in (?)", structIds)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

// DeleteByAdminId 删除某个管理员的组织权限节点
func (d *AdminStructDao) DeleteByAdminId(adminId string) error {
	if adminId == "" {
		return nil
	}
	_, err := core.NewDao(d.Model).Delete("`admin_id` = ?", adminId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByStructId 删除某个组织的管理员
func (d *AdminStructDao) DeleteByStructId(structId string) error {
	if structId == "" {
		return nil
	}
	_, err := core.NewDao(d.Model).Delete("`struct_id` = ?", structId)
	if err != nil {
		return err
	}
	return nil
}

// InsertList 插入管理员组织节点
func (d *AdminStructDao) InsertList(adminId string, structIds []string) int64 {
	var affected int64 = 0
	for _, id := range structIds {
		if id == "" {
			continue
		}
		err := core.NewDao(d.Model).InsertNoId([]string{"admin_id", "struct_id"}, []any{adminId, id})
		if err == nil {
			affected++
		}
	}
	return affected
}
