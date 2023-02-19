package system

import (
	"fmt"
	"rtsys/utils/core"
)

type GenerateService struct {
	Db string
}

// TableList 获取表的列表
func (g *GenerateService) TableList() []string {
	tables := make([]string, 0)

	//show tables 只会返回一个表名字段的列表
	rows, err := core.G_DB.Conn(g.Db).Query("show tables")
	if err != nil {
		return tables
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err == nil {
			tables = append(tables, name)
		}
	}
	return tables
}

// TableExist 判断表是否存在
func (g *GenerateService) TableExist(table string) (exist bool) {
	exist = false
	if table == "" {
		return
	}
	rows, err := core.G_DB.Conn(g.Db).Query(fmt.Sprintf("show tables like '%s'", table))
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err == nil {
			if table == name {
				exist = true
				return
			}
		}
	}
	return
}

func (g *GenerateService) TableFields(table string) {

}
