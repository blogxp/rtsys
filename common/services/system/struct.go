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
	"rtsys/utils/tool"
	. "rtsys/utils/types"
	"strings"
	"sync"
)

type StructService struct {
	Dao *StructDao
}

var (
	instanceStructService *StructService //单例的对象
	onceStructService     sync.Once
)

func NewStructService() *StructService {
	onceStructService.Do(func() {
		instanceStructService = &StructService{Dao: NewStructDao()}
	})
	return instanceStructService
}

// TypeList 类型集合
func (s *StructService) TypeList() []KeyVal {
	return []KeyVal{{Key: "M", Value: "目录"}, {Key: "C", Value: "菜单"}, {Key: "F", Value: "按钮"}}
}

// TreeList 获取树形展示列表
// @params root 是否添加根节点
func (s *StructService) TreeList() []StructModel {
	list := s.Dao.MenuTreeList()
	return *list
}

// GetName 获取组织名称
func (s *StructService) GetName(id string, showParent bool) string {
	if id == "0" || id == "" {
		return ""
	}
	info := s.GetInfo(id)
	if info == nil {
		return ""
	}
	if !showParent {
		return info.Name
	}
	list := strings.Split(info.ParentName, ",")
	list = list[1:]
	list = append(list, info.Name)
	return strings.Join(list, "/")
}

func (s *StructService) GetInfo(id string) *StructModel {
	return s.Dao.GetInfoById(id)
}

// UpdateLevelsInfo 当修改组织构架时，修改子类所有的parent_name和levels
func (s *StructService) UpdateLevelsInfo(pid string) {
	if pid == "" || pid == "0" {
		return
	}
	parentInfo := s.GetInfo(pid)
	if parentInfo == nil {
		return
	}
	parentName := strings.Trim(parentInfo.ParentName+","+parentInfo.Name, ",")
	levels := strings.Trim(parentInfo.Levels+","+parentInfo.Id, ",")

	childList := s.Dao.GetListByParentId(pid)
	for _, item := range *childList {
		if item.ParentName != parentName || item.Levels != levels {
			effected, err2 := core.NewDao(s.Dao.Model).Update([]string{"parent_name", "levels"}, []any{parentName, levels}, "`id` = ?", item.Id)
			if err2 != nil {
				return
			}
			if effected > 0 {
				s.UpdateLevelsInfo(item.Id)
			}
		}
	}
}

// GetChildAllIdList 获取某个节点的所有子节点ID数组
func (s *StructService) GetChildAllIdList(pid string) []string {
	idList := make([]string, 0)
	lists := s.Dao.GetChildAllList(pid)
	for _, item := range *lists {
		idList = append(idList, item.Id)
	}
	return idList
}

// ReRootList 将一级的parent_id改为0
func (s *StructService) ReRootList(list []StructModel) []StructModel {
	ids := make([]string, len(list))
	newList := make([]StructModel, len(list))
	for i, v := range list {
		ids[i] = v.Id
	}
	for i, v := range list {
		if !tool.InArray(v.ParentId, ids) {
			v.ParentId = "0"
		}
		newList[i] = v
	}
	return newList
}
