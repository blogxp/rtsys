package tool

import (
	"fmt"
	"rtsys/config"
	"rtsys/utils/core"
)

type SchemaOperate struct {
	Tags   string                   //数据库连接标识符
	fields map[string][]SchemaField //表字段
}

// SchemaField 表字段信息结构体
type SchemaField struct {
	Name    string  //COLUMN_NAME 字段名
	Type    string  //DATA_TYPE 字段类型
	Key     string  //COLUMN_KEY 主键类型
	Comment *string //COLUMN_COMMENT 备注
	Max     *int    //CHARACTER_MAXIMUM_LENGTH 字段长度
}

func NewSchemaOperate(tags string) *SchemaOperate {
	return &SchemaOperate{Tags: tags, fields: map[string][]SchemaField{}}
}

// DbConfig 获取配置信息
func (sch *SchemaOperate) DbConfig() *config.DataBaseConf {
	if sch.Tags == "" {
		sch.Tags = "default"
	}
	val, ok := core.G_CONFIG.DataBase[sch.Tags]
	if ok {
		return &val
	}
	return nil
}

// TableList 获取表的列表
func (sch *SchemaOperate) TableList() []string {
	tables := make([]string, 0)
	//show tables 只会返回一个表名字段的列表
	rows, err := core.G_DB.Conn(sch.Tags).Query("show tables")
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
func (sch *SchemaOperate) TableExist(table string) (exist bool) {
	exist = false
	if table == "" {
		return
	}
	rows, err := core.G_DB.Conn(sch.Tags).Query(fmt.Sprintf("show tables like '%s'", table))
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

// TableFields 获取表的字段信息
func (sch *SchemaOperate) TableFields(table string) []SchemaField {
	if v, exist := sch.fields[table]; exist {
		return v
	}

	fields := make([]SchemaField, 0)

	conf := sch.DbConfig()
	if conf == nil {
		return fields
	}
	rows, err := core.G_DB.Conn(sch.Tags).Query(fmt.Sprintf("select COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT,CHARACTER_MAXIMUM_LENGTH,COLUMN_KEY from information_schema.COLUMNS where table_name='%s' and table_schema='%s'", table, conf.DbName))
	if err != nil {
		fmt.Println("查询表字段失败：", err)
		return fields
	}
	defer rows.Close()
	for rows.Next() {
		var field = SchemaField{}
		if err := rows.Scan(&field.Name, &field.Type, &field.Comment, &field.Max, &field.Key); err == nil {
			fields = append(fields, field)
		}
	}
	sch.fields[table] = fields
	return fields
}
