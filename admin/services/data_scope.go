// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package services

import (
	"fmt"
	"rtsys/common/services/system"
	"rtsys/utils/tool"
	"rtsys/utils/types"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// DataScopeTypeList 数据权限类型
func DataScopeTypeList() []types.KeyVal {
	return []types.KeyVal{{Key: "1", Value: "全部数据权限"}, {Key: "2", Value: "本部门及以下数据权限"}, {Key: "4", Value: "本部门数据权限"}, {Key: "8", Value: "自定数据权限"}, {Key: "16", Value: "仅本人数据权限"}}
}

type DataScopeFilter struct {
	LoginData      *LoginData
	powerData      *dataScopeUserPower
	powerDataLevel *dataScopeUserPower
	sync.Mutex
}

func NewDataScopeFilterByCtx(ctx *gin.Context) *DataScopeFilter {
	return &DataScopeFilter{powerData: nil, LoginData: GetLoginByCtx(ctx)}
}

type dataScopeUserPower struct {
	NoPower  bool // true 无任何权限
	AllPower bool // true 所有权限
	IsUser   bool //true代表自己的数据

	StructIds []string // 有权限的组织id数组(包含本部门，本部门及以下，自定义部门)，使用in查询

	//下面是将组织的levels字段存储在每条记录中，使用find_in_set和=的方式来查询
	//StructIds 在此时代表自定义数据权限的组织和本部门
	StructLevels string //本部门及以下
}

// GetQueryParams 生成查询where条件 和 占位替代数组
// 查询用户角色的数据权限，主要是查询所有授权的组织ID，然后拼接sql使用in
// 要求每个二级组织下面的所有子组织不能超过1000个（mysql默认in限制）
func (ds *DataScopeFilter) GetQueryParams(structField string, userField string) (whereStr string, args []any) {
	whereStr = " 0 "
	args = make([]any, 0)

	if structField == "" && userField == "" { //字段都存在则默认无权限
		return
	}
	powerData := ds.getPowerData(false)
	if powerData.NoPower { //无权限
		return
	}
	if powerData.AllPower { //所有权限
		whereStr = ""
		return
	}
	if len(powerData.StructIds) == 0 && !powerData.IsUser { //未查询到组织及非个人 无权限
		return
	}

	var (
		whereUser   string
		whereStruct string
	)
	if userField != "" {
		if powerData.IsUser {
			whereUser = " `" + userField + "` = ? "
			args = append(args, ds.LoginData.Id)
		}
	}
	if structField != "" {
		if len(powerData.StructIds) > 0 {
			if len(powerData.StructIds) == 1 {
				whereStruct = " `" + structField + "` = ? "
				args = append(args, powerData.StructIds[0])
			} else {
				whereStruct = " `" + structField + "` in (?) "
				args = append(args, powerData.StructIds)
			}
		}
	}
	if whereStruct == "" && whereUser == "" {
		return
	}
	if whereStruct != "" && whereUser != "" {
		whereStr = "(" + whereUser + " OR " + whereStruct + ")"
	} else {
		whereStr = whereUser + whereStruct
	}
	return
}

// CheckByFiled 根据传入的组织ID和用户ID 判断是否有权限
func (ds *DataScopeFilter) CheckByFiled(structId string, userId string) bool {
	if structId == "" && userId == "" {
		return false
	}
	powerData := ds.getPowerData(false)
	if powerData.NoPower { //无权限
		return false
	}
	if powerData.AllPower { //所有权限
		return true
	}
	if len(powerData.StructIds) == 0 && !powerData.IsUser { //未查询到组织及非个人 无权限
		return false
	}
	if structId != "" && tool.InArray(structId, powerData.StructIds) {
		return true
	}
	if userId != "" && powerData.IsUser && userId == ds.LoginData.Id {
		return true
	}
	return false
}

// getPowerData 获取根据parseDataScopePower解析的权限
func (ds *DataScopeFilter) getPowerData(isLevel bool) *dataScopeUserPower {
	ds.Lock()
	if isLevel {
		if ds.powerDataLevel == nil {
			ds.powerDataLevel = ds.parseDataScopePowerByLevels()
		}
	} else {
		if ds.powerData == nil {
			ds.powerData = ds.parseDataScopePower()
		}
	}
	ds.Unlock()
	if isLevel {
		return ds.powerDataLevel
	} else {
		return ds.powerData
	}

}

// parseDataScopePower 根据登录信息解析权限 生成dataScopeUserParse
func (ds *DataScopeFilter) parseDataScopePower() *dataScopeUserPower {
	res := &dataScopeUserPower{}

	if ds.LoginData == nil {
		res.NoPower = true //无权限
		return res
	}
	userId, _ := strconv.Atoi(ds.LoginData.Id)
	if userId < 1 {
		res.NoPower = true //无权限
		return res
	}
	if ds.LoginData.IsAdmin == "1" {
		res.AllPower = true //超管返回全部权限
		return res
	}

	dataScope, _ := strconv.Atoi(ds.LoginData.DataScope)
	if dataScope < 1 {
		res.NoPower = true //无权限
		return res
	}

	if (31 & dataScope) == 0 {
		res.NoPower = true //无效的键值
		return res
	}

	structIdInt, _ := strconv.Atoi(ds.LoginData.StructId)
	if structIdInt < 1 {
		res.NoPower = true //无组织 返回无权限
		return res
	}

	if len(ds.LoginData.RoleList) < 1 {
		res.NoPower = true //无角色返回无权限
		return res
	}

	//权限范围的组织ID
	structIds := make([]string, 0)

	if 1&dataScope == 1 { //全部数据权限
		res.AllPower = true
		return res
	}

	if 2&dataScope == 2 { //本部门及以下数据权限
		structIds = append(structIds, ds.LoginData.StructId)
		childList := system.NewStructService().GetChildAllIdList(ds.LoginData.StructId)
		if len(childList) > 0 {
			structIds = append(structIds, childList...)
		}
	}

	if 4&dataScope == 4 { //本部门数据权限
		structIds = append(structIds, ds.LoginData.StructId)
	}

	if 8&dataScope == 8 { //自定义
		for _, role := range ds.LoginData.RoleList {
			roleStructs := system.NewRoleStructService().GetRoleStructIdList(role)
			if len(roleStructs) > 0 {
				structIds = append(structIds, roleStructs...)
			}
		}
	}
	if 16&dataScope == 16 { //个人数据
		res.IsUser = true
	}

	structIds = tool.UniqueArrStr(structIds)
	res.StructIds = structIds

	return res
}

//////      信息表中增加融入字段 struct_levels存储组织的上级树，当单独二级以下的子组织较多时将近1000或大于1000   /////////

// GetQueryParamsByLevel 基于组织架构的levels 来判断子组织
func (ds *DataScopeFilter) GetQueryParamsByLevel(structField string, structLevelField string, userField string) (whereStr string, args []any) {
	whereStr = " 0 "
	args = make([]any, 0)

	if structLevelField == "" || structField == "" {
		return
	}

	powerData := ds.getPowerData(true)
	if powerData.NoPower { //无权限
		return
	}
	if powerData.AllPower { //所有权限
		whereStr = ""
		return
	}
	if len(powerData.StructIds) == 0 && !powerData.IsUser && powerData.StructLevels == "" { //未查询到组织及非个人 无权限
		return
	}

	var (
		whereStruct string
		whereUser   string
	)

	if userField != "" {
		if powerData.IsUser {
			whereUser = " `" + userField + "` = ? "
			args = append(args, ds.LoginData.Id)
		}
	}

	if len(powerData.StructIds) > 0 {
		if len(powerData.StructIds) == 1 {
			whereStruct = " `" + structField + "` = ? "
			args = append(args, powerData.StructIds[0])
		} else {
			whereStruct = " `" + structField + "` in (?) "
			args = append(args, powerData.StructIds)
		}
	}
	if powerData.StructLevels != "" {
		if whereStruct != "" {
			whereStruct += "OR"
		}
		whereStruct += " find_in_set ( ? ,`" + structLevelField + "`)"
		args = append(args, powerData.StructLevels)
	}
	if whereStruct == "" && whereUser == "" {
		return
	}
	if whereStruct != "" && whereUser != "" {
		whereStr = "(" + whereUser + " OR " + whereStruct + ")"
	} else {
		whereStr = whereUser + whereStruct
	}
	fmt.Println(whereUser)
	return
}

// CheckByFiledLevels 根据传入的组织ID和用户ID 判断是否有权限
func (ds *DataScopeFilter) CheckByFiledLevels(structId string, structLevel string, userId string) bool {
	if structId == "" || structLevel == "" {
		return false
	}
	powerData := ds.getPowerData(true)
	if powerData.NoPower { //无权限
		return false
	}
	if powerData.AllPower { //所有权限
		return true
	}
	if len(powerData.StructIds) == 0 && !powerData.IsUser && powerData.StructLevels == "" { //未查询到组织及非个人 无权限
		return false
	}

	if structLevel != "" && powerData.StructLevels != "" {
		if strings.Index(structLevel+",", powerData.StructLevels+",") > -1 {
			return true
		}
	}

	if structId != "" && tool.InArray(structId, powerData.StructIds) {
		return true
	}
	if userId != "" && powerData.IsUser && userId == ds.LoginData.Id {
		return true
	}
	return false
}

// parseDataScopePowerByLevels 使用组织结构的levels判断
func (ds *DataScopeFilter) parseDataScopePowerByLevels() *dataScopeUserPower {
	res := &dataScopeUserPower{}

	if ds.LoginData == nil {
		res.NoPower = true //无权限
		return res
	}
	userId, _ := strconv.Atoi(ds.LoginData.Id)
	if userId < 1 {
		res.NoPower = true //无权限
		return res
	}
	if ds.LoginData.IsAdmin == "1" {
		res.AllPower = true //超管返回全部权限
		return res
	}

	dataScope, _ := strconv.Atoi(ds.LoginData.DataScope)
	if dataScope < 1 {
		res.NoPower = true //无权限
		return res
	}

	if (31 & dataScope) == 0 {
		res.NoPower = true //无效的键值
		return res
	}

	structIdInt, _ := strconv.Atoi(ds.LoginData.StructId)
	if structIdInt < 1 {
		res.NoPower = true //无组织 返回无权限
		return res
	}

	if len(ds.LoginData.RoleList) < 1 {
		res.NoPower = true //无角色返回无权限
		return res
	}

	//权限范围的组织ID
	structIds := make([]string, 0)

	if 1&dataScope == 1 { //全部数据权限
		res.AllPower = true
		return res
	}

	if 2&dataScope == 2 { //本部门及以下数据权限
		res.StructLevels = ds.LoginData.StructId
		structIds = append(structIds, ds.LoginData.StructId)
	}

	if 4&dataScope == 4 { //本部门数据权限
		structIds = append(structIds, ds.LoginData.StructId)
	}

	if 8&dataScope == 8 { //自定义
		for _, role := range ds.LoginData.RoleList {
			roleStructs := system.NewRoleStructService().GetRoleStructIdList(role)
			if len(roleStructs) > 0 {
				structIds = append(structIds, roleStructs...)
			}
		}
	}
	if 16&dataScope == 16 { //个人数据
		res.IsUser = true
	}

	structIds = tool.UniqueArrStr(structIds)
	res.StructIds = structIds

	return res
}
