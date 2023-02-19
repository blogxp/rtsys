// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	. "rtsys/common/daos/system"
	. "rtsys/common/models/system"
	"rtsys/utils/core"
	. "rtsys/utils/types"
	"sync"
)

type MenuService struct {
	Dao *MenuDao
}

var (
	instanceMenuService *MenuService //单例的对象
	onceMenuService     sync.Once
)

func NewMenuService() *MenuService {
	onceMenuService.Do(func() {
		instanceMenuService = &MenuService{Dao: NewMenuDao()}
	})
	return instanceMenuService
}

// TypeList 类型集合
func (s *MenuService) TypeList() []KeyVal {
	return []KeyVal{{Key: "M", Value: "目录"}, {Key: "C", Value: "菜单"}, {Key: "F", Value: "按钮"}}
}

// TreeList 获取树形展示列表
// @params root 是否添加根节点
func (s *MenuService) TreeList(root string) []MenuModel {
	list := s.Dao.MenuTreeList()
	if root == "1" {
		lists := []MenuModel{{Id: "0", ParentId: "-1", Name: "顶级菜单"}}
		lists = append(lists, *list...)
		return lists
	}
	return *list
}

func (s *MenuService) GetParentName(parentId string) string {
	if parentId == "0" {
		return "顶级菜单"
	}
	info := s.Dao.Model.New()
	err := core.NewDao(s.Dao.Model).First(info, parentId)
	if err != nil {
		return "查询错误"
	}
	return info.Name
}

// CheckPerms 查询权限节点
func (s *MenuService) CheckPerms(perms string) string {
	if perms == "" {
		return ""
	}
	info := s.Dao.Model.New()
	err := core.NewDao(s.Dao.Model).SetField("id").First(info, "`perms` = ?", perms)
	if err != nil {
		return ""
	}
	return info.Id
}
