// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

////////       对dchest/captcha的进一步封装  支持local和redis           ///////
////////       NewCaptcha(store,conn) 两个可变参数 store:local和redis   conn 为使用哪个redis           ///////
////////       NewCaptcha().NewLen(xx)                ///////
////////       NewCaptcha().VerifyString(xx,xx)      ///////

package tool

import (
	"log"
	"rtsys/utils/core"
	"time"

	"github.com/dchest/captcha"
)

type CaptchaTool struct {
	Store   string
	Redis   string
	ExpTime time.Duration //redis有效
}

func NewCaptcha(args ...string) *CaptchaTool {
	store := "local"
	conn := "default" //redis的前缀
	if len(args) > 0 {
		store = args[0]
	}
	if len(args) > 1 {
		conn = args[1]
	}
	return &CaptchaTool{Store: store, ExpTime: time.Second * 10, Redis: conn}
}

func (c *CaptchaTool) NewLen(len int) string {
	if c.Store == "redis" {
		captcha.SetCustomStore(c)
	}
	return captcha.NewLen(len)
}
func (c *CaptchaTool) VerifyString(id string, digits string) bool {
	if c.Store == "redis" {
		captcha.SetCustomStore(c)
	}
	return captcha.VerifyString(id, digits)
}

func (c *CaptchaTool) Set(id string, digits []byte) {
	err := core.G_Redis.Conn(c.Redis).Set(id, string(digits), c.ExpTime).Err()
	if err != nil {
		log.Println("验证码Id保存失败:" + err.Error())
	}
	return
}
func (c *CaptchaTool) Get(id string, clear bool) (digits []byte) {
	result, err := core.G_Redis.Conn(c.Redis).Get(id).Result()
	if clear {
		core.G_Redis.Conn(c.Redis).Del(id)
	}
	if err != nil {
		return nil
	}
	return []byte(result)
}
