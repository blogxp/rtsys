// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package core

// ** 简单封装的基于IModel快速增删改查操作 ** //

import (
	"database/sql"
	"errors"
	"fmt"
	"rtsys/utils/types"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Dao struct {
	IModel
	TableName string //默认为TModel的Table()，可以自定义
	Fields    string
	OrderBy   []types.KeyVal
	Limit     string
}

// NewDao 创建一个操作类，第二个为可变参数 指定数据库源
func NewDao(model IModel) *Dao {
	return &Dao{IModel: model, Fields: "*", TableName: model.Table()}
}

////    提供基本的四种方法 Get Select Exec  ExecNamed    ////

func (dm *Dao) Get(dest any, query string, args ...any) error {
	query1, args1, err1 := dm.ParseIn(query, args...)
	if err1 != nil {
		return err1
	}
	dm.ShowSql(query1, args1...)
	err := G_DB.Conn(dm.DataBase()).Get(dest, query1, args1...)
	if err != nil {
		return err
	}
	return nil
}

func (dm *Dao) Select(dest any, query string, args ...any) error {
	query1, args1, err1 := dm.ParseIn(query, args...)
	if err1 != nil {
		return err1
	}
	dm.ShowSql(query1, args1...)
	err := G_DB.Conn(dm.DataBase()).Select(dest, query1, args1...)
	if err != nil {
		return err
	}
	return nil
}

func (dm *Dao) Exec(query string, args ...any) (sql.Result, error) {
	query1, args1, err1 := dm.ParseIn(query, args...)
	if err1 != nil {
		return nil, err1
	}
	dm.ShowSql(query1, args1...)
	prepare, err := G_DB.Conn(dm.DataBase()).Prepare(query1)
	if err != nil {
		return nil, err
	}
	return prepare.Exec(args1...)
}
func (dm *Dao) ExecNamed(query string, arg any) (sql.Result, error) {
	dm.ShowSql(query)
	prepare, err := G_DB.Conn(dm.DataBase()).PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	return prepare.Exec(arg)
}

// ParseIn 对语句中的in进行处理
func (dm *Dao) ParseIn(query string, args ...any) (string, []any, error) {
	queryLower := strings.ToLower(query)
	if strings.Index(queryLower, "in") != -1 {
		return sqlx.In(query, args...)
	}
	return query, args, nil
}

///////////////////////  下面为提供的快捷查询方法  ////////////////////////////////////

// GetInfoByField 根据某个字段获取一条记录
func (dm *Dao) GetInfoByField(field string, value string, fields string) (IModel, error) {
	info := dm.IModel.INew()
	err := dm.SetField(fields).First(info, "`"+field+"` = ?", value)
	return info, err
}

// First 获取单条信息
// where 为id 或者 条件 id=? AND status=?
// args 为 where中的?
// 当where为 条件时，至少有一个args
func (dm *Dao) First(dest any, where string, args ...any) error {
	sqlStr := "select " + dm.Fields + " from " + dm.TableName
	if where != "" && len(args) == 0 {
		args = append(args, where)
		where = "`id`= ?"
	}
	if where != "" {
		sqlStr += " where " + where
	}
	order := dm.buildOrder()
	if order != "" {
		sqlStr += order
	}
	if dm.Limit != "" {
		sqlStr += " limit " + dm.Limit
	}

	if dest == nil {
		dest = dm.IModel
	}
	return dm.Get(dest, sqlStr, args...)
}

// Lists 获取列表 没啥特殊
func (dm *Dao) Lists(dest any, where string, args ...any) error {
	sqlStr := "select " + dm.Fields + " from " + dm.TableName
	if where != "" {
		sqlStr += " where " + where
	}
	order := dm.buildOrder()
	if order != "" {
		sqlStr += order
	}
	if dm.Limit != "" {
		sqlStr += " limit " + dm.Limit
	}
	return dm.Select(dest, sqlStr, args...)
}

// Delete 删除
func (dm *Dao) Delete(where string, args ...any) (int64, error) {
	sqlStr := "delete from " + dm.TableName
	if len(args) == 0 {
		args = append(args, where)
		where = "id= ?"
	}
	if where != "" {
		sqlStr += " where " + where
	}
	exec, err := dm.Exec(sqlStr, args...)
	if err != nil {
		return 0, err
	}
	affected, _ := exec.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return affected, nil
}

// Insert 插入 传入一个字段数组  一个值的数组
func (dm *Dao) Insert(fields []string, values []any) (string, error) {
	id := G_GENID.New().String()
	fields = append(fields, "id")
	values = append(values, id)
	err := dm.InsertNoId(fields, values)
	if err != nil {
		return "", err
	}
	return id, nil
}
func (dm *Dao) InsertNoId(fields []string, values []any) error {
	err := dm.insertExec(fields, values)
	if err != nil {
		return err
	}
	return nil
}
func (dm *Dao) insertExec(fields []string, values []any) error {
	if len(fields) != len(values) {
		return errors.New("键值数量不一致")
	}
	places := make([]string, len(fields))
	for i := 0; i < len(fields); i++ {
		fields[i] = "`" + fields[i] + "`"
		places[i] = "?"
	}
	sqlStr := "insert into " + dm.TableName + " (" + strings.Join(fields, ",") + ") value (" + strings.Join(places, ",") + ")"

	exec, err1 := dm.Exec(sqlStr, values...)
	if err1 != nil {
		return err1
	}
	affected, err2 := exec.RowsAffected()
	if err2 != nil {
		return err2
	}
	if affected == 0 {
		return errors.New("插入失败，影响0行")
	}
	return nil
}

// InsertNamed 命名插入，传入一个键是表字段名的map
func (dm *Dao) InsertNamed(maps map[string]any) (string, error) {
	maps["id"] = G_GENID.New().String()
	err := dm.InsertNamedNoId(maps)
	if err != nil {
		return "", err
	}
	return maps["id"].(string), nil
}

// InsertNamedNoId 命名插入，传入一个键是表字段名的map
func (dm *Dao) InsertNamedNoId(maps map[string]any) error {
	err := dm.insertNamedExec(maps)
	if err != nil {
		return err
	}
	return nil
}

func (dm *Dao) insertNamedExec(maps map[string]any) error {
	var places, fields []string
	for key, _ := range maps {
		places = append(places, ":"+key)
		fields = append(fields, "`"+key+"`")
	}
	sqlStr := "insert into " + dm.TableName + " (" + strings.Join(fields, ",") + ") value (" + strings.Join(places, ",") + ")"
	exec, err1 := dm.ExecNamed(sqlStr, maps)
	if err1 != nil {
		return err1
	}
	affected, err2 := exec.RowsAffected()
	if err2 != nil {
		return err2
	}
	if affected == 0 {
		return errors.New("插入失败，影响0行")
	}
	return nil
}

// Update 插入 传入一个字段数组  一个值的数组
func (dm *Dao) Update(fields []string, values []any, where string, args ...any) (int64, error) {
	if len(fields) < 1 {
		return 0, errors.New("无更新的字段")
	}
	if len(fields) != len(values) {
		return 0, errors.New("键值数量不一致")
	}
	places := make([]any, len(fields))
	for i := 0; i < len(fields); i++ {
		fields[i] = "`" + fields[i] + "` = ?"
		places[i] = values[i]
	}
	sqlStr := "update " + dm.TableName + " set " + strings.Join(fields, ",")
	if where != "" {
		if len(args) == 0 {
			places = append(places, where)
			where = "`id` = ?"
		} else {
			for j := 0; j < len(args); j++ {
				places = append(places, args[j])
			}
		}
		sqlStr += " where " + where
	}
	exec, err1 := dm.Exec(sqlStr, places...)
	if err1 != nil {
		return 0, err1
	}
	effected, err2 := exec.RowsAffected()
	if err2 != nil {
		return 0, nil
	}
	return effected, nil
}

// UpdateNamed 命名更新，传入一个键是表字段名的map
func (dm *Dao) UpdateNamed(maps map[string]any, where map[string]any) (int64, error) {
	var fields, whereArr []string
	for key, _ := range maps {
		fields = append(fields, "`"+key+"` = :"+key)
	}
	sqlStr := "update " + dm.TableName + " set " + strings.Join(fields, ",")
	if where != nil {
		for wkey, wval := range where {
			whereArr = append(whereArr, "`"+wkey+"` = :w_"+wkey)
			maps["w_"+wkey] = wval
		}
		sqlStr += " where " + strings.Join(whereArr, " AND ")
	}
	exec, err1 := dm.ExecNamed(sqlStr, maps)
	if err1 != nil {
		return 0, err1
	}
	effected, err2 := exec.RowsAffected()
	if err2 != nil {
		return 0, nil
	}
	return effected, nil
}

/////    对属性的复制处理        /////

// SetField 用于查询时定义返回的字段
func (dm *Dao) SetField(fields string) *Dao {
	if fields == "" {
		fields = "*"
	}
	dm.Fields = fields
	return dm
}
func (dm *Dao) SetOrderBy(order []types.KeyVal) *Dao {
	dm.OrderBy = order
	return dm
}
func (dm *Dao) SetLimit(limit string) *Dao {
	dm.Limit = limit
	return dm
}

func (dm *Dao) SetTable(table string) *Dao {
	dm.TableName = table
	return dm
}

func (dm *Dao) buildOrder() string {
	order := ""
	for _, val := range dm.OrderBy {
		field := val.Key
		if field == "" {
			continue
		}
		asc := val.Value
		if asc == "" {
			asc = "asc"
		}
		if order == "" {
			order += field + " " + asc
		} else {
			order += "," + field + " " + asc
		}
	}
	if order != "" {
		order = " order by " + order
	}
	return order
}

func (dm *Dao) ShowSql(sql string, args ...any) {
	if G_CONFIG.Server.DbShowSql {
		fmt.Printf(strings.Replace(sql+"\n", "?", "%v", -1), args...)
	}
}
