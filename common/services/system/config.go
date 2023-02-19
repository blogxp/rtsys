// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	. "rtsys/common/daos/system"
	"rtsys/utils/core"
	. "rtsys/utils/trans"
	. "rtsys/utils/types"
	"sync"
	"time"
)

type ConfigService struct {
	Dao *ConfigDao
}

var (
	instanceConfigService *ConfigService //单例的对象
	onceConfigService     sync.Once
)

func NewConfigService() *ConfigService {
	onceConfigService.Do(func() {
		instanceConfigService = &ConfigService{Dao: NewConfigDao()}
	})
	return instanceConfigService
}

// StyleList 配置类型集合
func (s *ConfigService) StyleList() map[string]string {
	return map[string]string{"text": "文本", "textarea": "多行文本", "array": "数组", "select": "枚举"}
}

func (s *ConfigService) GetArrayOne(code string) []KeyVal {
	var list []KeyVal
	info := s.Dao.GetInfoByType(code)
	if info == nil || info.Style != "array" {
		return list
	}
	list = StrToKeyValByLine(info.Value, "\n", ":")
	return list
}

func (s *ConfigService) GetValue(code string) string {
	info := s.Dao.GetInfoByType(code)
	if info == nil {
		return ""
	}
	return info.Value
}

// GroupList 获取配置分组并解析
func (s *ConfigService) GroupList() []KeyVal {
	return s.GetArrayOne("sys_config_group")
}

// GetStyleName 配置分组列表
func (s *ConfigService) GetStyleName(style string) string {
	name := ""
	if style != "" {
		styleList := s.StyleList()
		if item, ok := styleList[style]; ok {
			name = item
		}
	}
	return name
}

// StructTypeList 获取组织架构类型
func (s *ConfigService) StructTypeList() []KeyVal {
	return s.GetArrayOne("sys_struct_type")
}

// GetStructTypeName 获取组织类型名称
func (s *ConfigService) GetStructTypeName(code string) string {
	name := ""
	if code != "" {
		typeList := s.StructTypeList()
		for _, item := range typeList {
			if item.Key == code {
				name = item.Value
				break
			}
		}
	}
	return name
}

// GetGroupName 获取配置分组名称
func (s *ConfigService) GetGroupName(group string, groupList []KeyVal) string {
	name := ""
	if group != "" {
		if groupList == nil {
			groupList = s.GroupList()
		}
		for _, item := range groupList {
			if item.Key == group {
				name = item.Value
				break
			}
		}
	}
	return name
}

// GetListByGroup 根据分组获取配置列表
func (s *ConfigService) GetListByGroup() []map[string]any {
	var reList []map[string]any
	groupList := s.GroupList()
	if len(groupList) == 0 {
		return reList
	}
	list := s.Dao.GroupChildList()
	for _, group := range groupList {
		var child []map[string]any
		for _, item := range *list {
			if item.Groups == group.Key {
				var extra []KeyVal
				if item.Style == "select" {
					extra = StrToKeyValByLine(item.Extra, "\n", ":")
				}
				child = append(child, map[string]any{"id": item.Id, "title": item.Title, "style": item.Style, "note": item.Note, "value": item.Value, "extra": extra})
			}
		}
		reList = append(reList, map[string]any{"name": group.Value, "key": group.Key, "child": child})
	}
	return reList
}

func (s *ConfigService) UpdateByAutoKey(autokey *AutoKey) {
	dao := core.NewDao(s.Dao.Model)
	for id, val := range autokey.Where {
		_, _ = dao.Update([]string{"value", "update_time"}, []any{val, TimeFormat(time.Now(), "")}, "`id` = ?", id)
	}
}
