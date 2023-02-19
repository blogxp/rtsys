// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package lib

// SetStatusData 用来后台公共操作设置状态
type SetStatusData struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
	Field  string `json:"field"`
}

// RoleMenuDataAuth 角色菜单和数据权限提交绑定的结构体
type RoleMenuDataAuth struct {
	Id        string `json:"id"`
	TreeId    string `json:"tree_id"`
	DataScope string `json:"data_scope"`
}
