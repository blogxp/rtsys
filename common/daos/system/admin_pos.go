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

type AdminPosDao struct {
	Model *AdminPosModel
}

var (
	instanceAdminPosDao *AdminPosDao //单例的对象
	onceAdminPosDao     sync.Once
)

func NewAdminPosDao() *AdminPosDao {
	onceAdminPosDao.Do(func() {
		instanceAdminPosDao = &AdminPosDao{Model: NewAdminPosModel()}
	})
	return instanceAdminPosDao
}

// GetAdminPosList 获取某个管理员的岗位列表
func (d *AdminPosDao) GetAdminPosList(adminId string) (*[]AdminPosModel, error) {
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

// DeleteByAdminId 删除某个岗位的管理员
func (d *AdminPosDao) DeleteByAdminId(adminId string) error {
	if adminId == "" {
		return nil
	}
	_, err := core.NewDao(d.Model).Delete("`admin_id` = ?", adminId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByPosId 删除某个岗位的管理员
func (d *AdminPosDao) DeleteByPosId(posId string) error {
	if posId == "" {
		return nil
	}
	_, err := core.NewDao(d.Model).Delete("`pos_id` = ?", posId)
	if err != nil {
		return err
	}
	return nil
}

// InsertList 插入角色岗位节点
func (d *AdminPosDao) InsertList(adminId string, posIds []string) int64 {
	var affected int64 = 0
	for _, id := range posIds {
		if id == "" {
			continue
		}
		err := core.NewDao(d.Model).InsertNoId([]string{"admin_id", "pos_id"}, []any{adminId, id})
		if err == nil {
			affected++
		}
	}
	return affected
}
