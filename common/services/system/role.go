// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"errors"
	"fmt"
	"rtsys/admin/lib"
	. "rtsys/common/daos/system"
	. "rtsys/common/models/system"
	"rtsys/utils/core"
	"sync"
)

type RoleService struct {
	Dao *RoleDao
}

var (
	instanceRoleService *RoleService //单例的对象
	onceRoleService     sync.Once
)

func NewRoleService() *RoleService {
	onceRoleService.Do(func() {
		instanceRoleService = &RoleService{Dao: NewRoleDao()}
	})
	return instanceRoleService
}

// SaveMenuAuth 菜单权限保存
func (s *RoleService) SaveMenuAuth(form *lib.RoleMenuDataAuth) error {
	if form.Id == "" {
		return errors.New("角色参数错误")
	}
	roleInfo := s.Dao.GetInfoById(form.Id)
	if roleInfo == nil {
		return errors.New("角色信息错误")
	}
	err := NewRoleMenuService().UpdateRoleMenu(form.Id, form.TreeId)
	if err != nil {
		return err
	}
	return nil
}

// SaveDataScopeAuth 数据权限保存
func (s *RoleService) SaveDataScopeAuth(form *lib.RoleMenuDataAuth) error {
	if form.Id == "" {
		return errors.New("角色参数错误")
	}
	if form.DataScope == "" {
		return errors.New("数据范围错误")
	}
	roleInfo := s.Dao.GetInfoById(form.Id)
	if roleInfo == nil {
		return errors.New("角色信息错误")
	}
	_, err := core.NewDao(s.Dao.Model).Update([]string{"data_scope"}, []any{form.DataScope}, "`id` = ?", form.Id)
	if err != nil {
		fmt.Println(err)
		return errors.New("数据范围更新失败")
	}
	if form.DataScope != "8" {
		form.TreeId = ""
	}
	err1 := NewRoleStructService().UpdateRoleStruct(form.Id, form.TreeId)
	if err1 != nil {
		return err1
	}
	return nil
}

func (s *RoleService) GetInfo(id string) *RoleModel {
	return s.Dao.GetInfoById(id)
}

// GetName 获取角色名称
func (s *RoleService) GetName(id string) string {
	if id == "0" || id == "" {
		return ""
	}
	info := s.GetInfo(id)
	if info == nil {
		return ""
	}
	return info.Name
}
