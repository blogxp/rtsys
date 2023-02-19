// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package config

type Config struct {
	System   SystemConf              `yaml:"system"`
	Server   ServerConf              `yaml:"server"`
	Route    RouteConf               `yaml:"route"`
	DataBase map[string]DataBaseConf `yaml:"database"`
	Redis    map[string]RedisConf    `yaml:"redis"`
	Oss      OssConf                 `yaml:"oss"`
}
