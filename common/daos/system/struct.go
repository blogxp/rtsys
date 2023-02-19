// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"fmt"
	. "rtsys/common/models/system"
	"rtsys/utils/core"
	"rtsys/utils/types"
	"sync"
)

type StructDao struct {
	Model *StructModel
}

var (
	instanceStructDao *StructDao //单例的对象
	onceStructDao     sync.Once
)

func NewStructDao() *StructDao {
	onceStructDao.Do(func() {
		instanceStructDao = &StructDao{Model: NewStructModel()}
	})
	return instanceStructDao
}

func (d *StructDao) MenuTreeList() *[]StructModel {
	list := d.Model.NewSlice()
	_ = core.NewDao(d.Model).SetField("id,parent_id,name").SetOrderBy([]types.KeyVal{{Key: "parent_id"}, {Key: "list_sort"}, {Key: "id"}}).Lists(list, "")
	return list
}

func (d *StructDao) GetInfoById(id string) *StructModel {
	model := d.Model.New()
	err := core.NewDao(d.Model).First(model, id)
	if err != nil || model.Id == "" {
		return nil
	}
	return model
}

func (d *StructDao) GetListByParentId(parentId string) *[]StructModel {
	childList := d.Model.NewSlice()
	_ = core.NewDao(d.Model).Lists(childList, "`parent_id` = ?", parentId)
	return childList
}

func (d *StructDao) GetChildAllList(parentId string) *[]StructModel {
	lists := d.Model.NewSlice()
	sqlStr := "select * from " + d.Model.Table() + " where FIND_IN_SET(?,`levels`)"
	err := core.NewDao(d.Model).Select(lists, sqlStr, parentId)
	if err != nil {
		fmt.Println(err)
	}
	return lists
}
