// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package core

// ** 数据库初始化 ** //

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type B5Db struct {
	List map[string]*sqlx.DB
}

// InitDb 数据库配置初始化
// 将多数据库进行sqlx.Open打开并赋值给全局的G_DB
func InitDb() {
	db := &B5Db{}
	db.parseItem()
	G_DB = db
}

// Conn 数据库操作方法，返回sqlx.DB
func (db *B5Db) Conn(args ...string) *sqlx.DB {
	types := "default"
	if len(args) > 0 && args[0] != "" {
		types = args[0]
	}

	if val, exist := db.List[types]; exist {
		err := val.Ping()
		if err != nil {
			log.Println("数据库连接失败："+err.Error())
		}
		return val
	}
	log.Println("数据库配置错误：未找到" + types)
	return &sqlx.DB{}
}

// parseItem 解析并创建连接
func (db *B5Db) parseItem() {
	list := make(map[string]*sqlx.DB)
	for key, item := range G_CONFIG.DataBase {
		if item.Driver == "mysql" {
			list[key] = InitMysql(item)
		}
	}
	db.List = list
}
