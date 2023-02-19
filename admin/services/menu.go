// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package services

import (
	"rtsys/admin/lib"
	. "rtsys/common/daos/system"
	. "rtsys/common/models/system"
	"rtsys/utils/trans"

	"github.com/gin-gonic/gin"
)

func AdminMenuShowList(ctx *gin.Context) []map[string]any {
	loginData, err := CheckLoginCookie(ctx)
	if err != nil {
		return nil
	}
	menuDao := NewMenuDao()
	menuIdList := loginData.MenuList
	if loginData.IsAdmin == "1" {
		menuIdList = nil
	}
	menuList := menuDao.GetMenuShowLists(menuIdList)
	return AdminMenuTree(menuList, "0", 0)
}

func AdminMenuTree(menuList *[]MenuModel, pid string, deep int) []map[string]any {
	list := make([]map[string]any, 0)
	for _, item := range *menuList {
		if item.ParentId == pid && item.Status == "1" {
			row := trans.StructToMap(item, true)
			row["deep"] = deep
			childList := AdminMenuTree(menuList, item.Id, deep+1)
			if len(childList) < 1 {
				row["child"] = nil
			} else {
				row["child"] = childList
			}
			list = append(list, row)
		}
	}
	return list
}

func AdminMenuToHtml(list []map[string]any, deep int) string {
	html := ""
	if len(list) < 1 || list == nil {
		return html
	}
	for _, item := range list {

		if item["deep"] != deep {
			continue
		}
		name := item["name"].(string)
		icon := item["icon"].(string)
		if item["type"] == "C" {
			target := "menuItem"
			isRefresh := "false"
			url := lib.AdminCreateUrl(item["url"].(string))

			if item["target"] == "1" {
				target = "menuBlank"
			}
			if item["is_refresh"] == "1" {
				isRefresh = "true"
			}
			if item["parent_id"] == "0" {
				html += "<li><a class='" + target + "' href='" + url + "' data-refresh='" + isRefresh + "'>"
				if icon != "" {
					html += "<i class='" + icon + "'></i>"
				}
				html += " <span class='nav-label'>" + name + "</span></a></li>"
			} else {
				html += "<li><a class='" + target + "' href='" + url + "' data-refresh='" + isRefresh + "'>" + name + "</a></li>"
			}
		} else {
			child := trans.AnyToMapList(item["child"])
			if child == nil || deep >= 3 {
				continue
			}
			html += "<li><a href='javascript:;'>"
			if icon != "" {
				html += "<i class='" + icon + "'></i>"
			}
			html += " <span class='nav-label'>" + name + "</span><span class='fa arrow'></span></a>"
			deepClass := "nav-second-level"
			if deep > 0 {
				deepClass = "nav-third-level"
			}
			html += "<ul class='nav " + deepClass + "'>"
			html += AdminMenuToHtml(child, deep+1)
			html += "</ul></li>"
		}
	}
	return html
}
