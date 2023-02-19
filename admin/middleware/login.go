// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package middleware

import (
	"net/http"
	"rtsys/admin/controller/common"
	"rtsys/admin/lib"
	"rtsys/admin/services"
	"rtsys/utils/tool"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoginAdminMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		publicC := common.NewPublicController()
		publicUrl := publicC.ParseUrl("", true)
		path := ctx.Request.URL.Path
		if strings.Index(path, publicUrl) == 0 {
			return
		}
		data, err := services.CheckLoginCookie(ctx)
		if err != nil || data.Id == "" {
			loginUrl := publicC.ParseUrl("login", true)
			if tool.IsRender(ctx) {
				ctx.Redirect(http.StatusFound, loginUrl)
			} else {
				publicC.Error(ctx, "请先登录", nil, lib.RNoLogin, loginUrl)
			}
			ctx.Abort()
			return
		} else {
			//往中间件中设置登录信息
			ctx.Set("_login_", data)
		}
	}
}
