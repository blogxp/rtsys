// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

///////   示例 登录  ///////

package store

import (
	"rtsys/api/lib"
	"rtsys/common/daos/system"

	"github.com/gin-gonic/gin"
)

type PublicApiStore struct {
	lib.BaseApi
}

func NewPublicApiStore() *PublicApiStore {
	c := &PublicApiStore{}
	c.Id = "public"
	return c
}

func (c *PublicApiStore) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("login", c.Login))
}

// Login 生成登录token
func (c *PublicApiStore) Login(ctx *gin.Context) {
	token, err := system.NewAppTokenDao().SetToken("1001", "store", "app", "")
	if err != nil {
		c.Error(ctx, err.Error())
		return
	}
	c.Success(ctx, "登录成功", gin.H{"token": token})
}
