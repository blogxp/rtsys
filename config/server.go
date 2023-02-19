// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package config

type ServerConf struct {
	Domain    string `yaml:"domain"`      //外网访问域名
	OssDomain string `yaml:"oss-domain"`  //文件外网访问域名，用于文件域名拼接等
	Port      string `yaml:"port"`        //服务器端口
	OssType   string `yaml:"oss-type"`    //文件上传保存方式
	WorkerId  int64  `yaml:"worker-id"`   //分布式雪花算法中的机器码
	DeBug     bool   `yaml:"debug"`       //开发模式 1debug模式  0 正式模式
	NetWriter bool   `yaml:"net-writer"`  //是否控制器输除网络请求
	DbShowSql bool   `yaml:"db-show-sql"` //是否打印sql
}
