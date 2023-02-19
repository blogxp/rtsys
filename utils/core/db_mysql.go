// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package core

import (
	"log"
	"rtsys/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitMysql(conf config.DataBaseConf) *sqlx.DB {
	dsn := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.DbName + "?charset=" + conf.Charset + "&parseTime=true"
	DbConn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Printf("%s数据库连接失败：%s\n", conf.Type, err.Error())
	}
	DbConn.SetMaxOpenConns(500) // 最大连接数
	DbConn.SetMaxIdleConns(20)  // 最大空闲连接数
	return DbConn
}
