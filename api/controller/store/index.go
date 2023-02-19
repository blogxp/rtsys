// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

/////////  示例  当前控制器内有需要登录判断的方法和不需登录判断的方法    ///////////

package store

import (
	"fmt"
	"rtsys/api/services"
	"rtsys/utils/core"

	"github.com/gin-gonic/gin"
)

// IndexApiStore  嵌套LoginCheck结构体
type IndexApiStore struct {
	LoginCheck
}

func (c *IndexApiStore) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", c.Index))
	group.POST(c.Dispatch("user", c.User))
}

func NewIndexApiStore() *IndexApiStore {
	c := &IndexApiStore{}
	c.Id = "index"

	//对LoginCheck结构处理
	c.LoginCheck.NoLoginActions = []string{"index"}

	//定义BeforeAction
	c.BaseApi.BeforeAction = c.LoginCheck.Handle
	return c
}

func (c *IndexApiStore) Index(ctx *gin.Context) {
	err := core.G_Redis.Conn().Set("asdad", "2132131", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	c.Success(ctx, "index")
}

func (c *IndexApiStore) User(ctx *gin.Context) {
	//获取登录信息
	appToken := services.GetApiLoginInfo(ctx)
	fmt.Println(appToken)
	c.Success(ctx, "user", appToken)
}
