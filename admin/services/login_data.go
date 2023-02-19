// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package services

import (
	"errors"
	"rtsys/utils/tool"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type LoginData struct {
	Id        string
	NickName  string
	DataScope string
	RoleList  []string
	MenuList  []string
	IsAdmin   string
	StructId  string
}

// CheckLoginCookie 检测获取登录信息
func CheckLoginCookie(ctx *gin.Context) (*LoginData, error) {
	if cookie, err := ctx.Cookie("b5gocmf_login"); err == nil {
		data, err1 := ParseLoginCookie(cookie)
		if err1 != nil {
			RemoveLoginCookie(ctx)
			return nil, err1
		}
		return data, nil
	}
	return nil, errors.New("登录失效：已过期")
}

func GetLoginByCtx(ctx *gin.Context) *LoginData {
	loginCache, exists := ctx.Get("_login_")
	if !exists {
		return nil
	}
	switch loginCache.(type) {
	case *LoginData:
		return loginCache.(*LoginData)
	}
	return nil
}

// ParseLoginCookie 解析登录cookie
func ParseLoginCookie(cookie string) (*LoginData, error) {
	claims, err := tool.ParseTokenJwt(cookie)
	if err != nil {
		return nil, err
	}
	data, ok := claims["data"]
	if !ok {
		return nil, errors.New("登录失效:信息丢失")
	}
	marshal, errJm := json.Marshal(data)
	if errJm != nil {
		return nil, errors.New("登录失效:序列化失败")
	}
	login := &LoginData{}
	errJu := json.Unmarshal(marshal, login)
	if errJu != nil {
		return nil, errors.New("登录失效:解析失败")
	}
	if login.Id == "" {
		return nil, errors.New("登录失效:请重新登录")
	}
	return login, nil
}

// CreateLoginCookie 创建登录cookie标识
func CreateLoginCookie(data *LoginData, exp int) (string, error) {
	if exp <= 0 {
		exp = 24 * 60 * 60
	}
	exps := time.Now().Unix() + int64(exp)
	token, err := tool.CreateTokenJwt(data, exps)
	if err != nil {
		return "", err
	}
	return token, nil
}

// RemoveLoginCookie 删除登录cookie
func RemoveLoginCookie(ctx *gin.Context) {
	ctx.SetCookie("b5gocmf_login", "", -1, "/", "", false, true)
}
