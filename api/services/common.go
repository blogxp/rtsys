// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package services

import (
	. "rtsys/common/models/system"

	"github.com/gin-gonic/gin"
)

// GetApiLoginInfo 获取登录信息
func GetApiLoginInfo(ctx *gin.Context) *AppTokenModel {
	appToken, exists := ctx.Get("_token_")
	if !exists {
		return nil
	}
	switch appToken.(type) {
	case *AppTokenModel:
		model := appToken.(*AppTokenModel)
		if model.UserId == "" {
			return nil
		}
		return model
	}
	return nil
}
