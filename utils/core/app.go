// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

func Load() *gin.Engine {
	//加载配置文件
	loadConfig("./")

	//加载雪花算法
	loadSnowFlake()

	//初始化validator
	LoadValidator()

	//连接数据库
	InitDb()

	//连接Redis
	InitRedis()

	//设置模式
	if !G_CONFIG.Server.DeBug {
		gin.SetMode(gin.ReleaseMode)
	}
	//打印网络请求
	if !G_CONFIG.Server.NetWriter {
		gin.DefaultWriter = io.Discard
	}
	engine := gin.New()

	//使用自定义格式的日志log
	engine.Use(gin.LoggerWithFormatter(CustomB5Log)).Use(gin.Recovery())

	return engine
}

// LoadConfig 加载解析config.yaml 配置文件，并赋值给全局配置utils.CONIFG
func loadConfig(path string) {
	file, err := os.ReadFile(path + "config.yaml")
	if err != nil {
		log.Fatal("配置文件读取失败：", err)
	}

	err = yaml.Unmarshal(file, &G_CONFIG)
	if err != nil {
		log.Fatal("配置文件解析失败：", err)
		return
	}
}

func loadSnowFlake() {
	node, err := NewNode(G_CONFIG.Server.WorkerId)
	if err != nil {
		log.Fatal("分布式ID配置加载失败：", err)
		return
	}
	G_GENID = node
}

func Run(engine *gin.Engine) {
	fmt.Printf(`
###########################################################################
// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author  : 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------
// | WebSite : www.b5net.com
// +----------------------------------------------------------------------
###########################################################################
`)
	if err := engine.Run(":" + G_CONFIG.Server.Port); err != nil {
		log.Fatal("服务启动失败：", err)
	}
}

// TestLoad 用于测试时 加载全局配置
func TestLoad(path string) {
	//加载配置文件
	loadConfig(path)
	//加载雪花算法
	loadSnowFlake()

	//初始化validator
	LoadValidator()

	//连接数据库
	InitDb()

	//连接Redis
	InitRedis()

}
