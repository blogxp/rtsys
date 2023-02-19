// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package common

import (
	"net/http"
	"os"
	"path/filepath"
	"rtsys/admin/lib"
	"rtsys/admin/services"
	. "rtsys/common/daos/system"
	. "rtsys/common/services/system"
	"rtsys/utils/tool"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

type PublicController struct {
	lib.Controller
}

func NewPublicController() *PublicController {
	c := &PublicController{}
	c.Id = "public"
	return c
}

func (c *PublicController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("login", false, c.Login))
	group.POST(c.Dispatch("tologin", false, c.ToLogin))
	group.GET(c.Dispatch("logout", false, c.LogOut))
	group.POST(c.Dispatch("captcha_id", false, c.CaptchaId))
	group.GET(c.Dispatch("captcha_img", false, c.CaptchaImg))
	group.GET(c.Dispatch("download", false, c.DownLoad))
	group.POST(c.Dispatch("test", false, c.Test))
	group.GET(c.Dispatch("test1", false, c.Test1))
}

func (c *PublicController) Test(ctx *gin.Context) {
	service := NewConfigService()
	c.Success(ctx, "", map[string]string{"appid": service.GetValue("wechat_appid"), "appsecret": service.GetValue("wechat_appsecret")})
}
func (c *PublicController) Test1(ctx *gin.Context) {
	c.Render(ctx, "login", nil)
}

func (c *PublicController) Login(ctx *gin.Context) {
	_, err := services.CheckLoginCookie(ctx)
	if err == nil {
		ctx.Redirect(http.StatusMovedPermanently, NewIndexController().ParseUrl("index", true))
		return
	}
	c.Render(ctx, "login", nil)
}

func (c *PublicController) ToLogin(ctx *gin.Context) {
	form := &services.LoginForm{}
	err := ctx.ShouldBind(form)
	if err != nil {
		c.Error(ctx, "表单绑定失败")
		return
	}
	err1 := form.Login(ctx)

	status := "1"
	msg := "登录成功"
	if err1 != nil {
		status = "0"
		msg = err1.Error()
	}
	NewLoginLogDao().AddLog(ctx, form.UserName, status, msg)

	if err1 != nil {
		c.Error(ctx, err1.Error())
		return
	}
	c.Success(ctx, "登录成功")
}

func (c *PublicController) LogOut(ctx *gin.Context) {
	services.RemoveLoginCookie(ctx)
	c.Success(ctx, "成功退出")
}

func (c *PublicController) CaptchaId(ctx *gin.Context) {
	id := tool.NewCaptcha("local").NewLen(4)
	c.Success(ctx, "", gin.H{"id": id})
}
func (c *PublicController) CaptchaImg(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" || !captcha.Reload(id) {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.Request.Header.Set("Content-Type", "image/png")
	_ = captcha.WriteImage(ctx.Writer, id, 130, 40)
}

func (c *PublicController) DownLoad(ctx *gin.Context) {
	url := ctx.DefaultQuery("fileName", "")
	if url == "" {
		c.Error(ctx, "文件参数错误")
		return
	}
	del := ctx.DefaultQuery("delete", "")
	if del == "true" {
		defer func() {
			_ = os.Remove(filepath.Clean("." + url))
		}()
	}
	ctx.File(filepath.Clean("." + url))
}
