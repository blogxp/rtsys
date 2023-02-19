// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package services

import (
	"crypto/md5"
	"errors"
	"fmt"
	. "rtsys/common/models/system"
	. "rtsys/common/services/system"
	"rtsys/utils/core"
	"rtsys/utils/tool"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	UserName   string `form:"username" validate:"required" label:"账号"`
	Password   string `form:"password" validate:"required" label:"密码"`
	CaptchaId  string `form:"captcha_id" validate:"required" label:"验证码标识"`
	CaptchaImg string `form:"captcha_img" validate:"required" label:"验证码"`
	Remember   string `form:"remember"`
}

// Login 登录处理
func (l *LoginForm) Login(ctx *gin.Context) error {
	admin, err := l.Check()
	if err != nil {
		return err
	}
	maxAge := 0
	if l.Remember == "1" {
		maxAge = 72 * 3600
	}
	//cookie保存数据
	loginData := &LoginData{Id: admin.Id, NickName: admin.RealName, DataScope: "0", IsAdmin: "0"}

	if admin.Id == core.G_ADMIN_ID {
		loginData.IsAdmin = "1"
	}

	//获取角色和数据权限
	dataScope := 0
	roleList := make([]string, 0)
	if loginData.IsAdmin != "1" {
		roles := NewAdminRoleService().GetRoleIdListByAdmin(admin.Id)
		for _, roleId := range roles {
			if roleId == core.G_ROLE_ID {
				loginData.IsAdmin = "1"
				dataScope = 0
				break
			} else {
				roleInfo := NewRoleService().GetInfo(roleId)
				if roleInfo != nil {
					if atoi, errDs := strconv.Atoi(roleInfo.DataScope); errDs == nil {
						dataScope += atoi
					}
					roleList = append(roleList, roleId)
				}
			}
		}
		loginData.RoleList = roleList
	}

	if loginData.IsAdmin != "1" && len(roleList) < 1 {
		return errors.New("无角色分组，登录失败")
	}

	//获取菜单列表
	if loginData.IsAdmin != "1" {
		loginData.DataScope = strconv.Itoa(dataScope)
		menuList := make([]string, 0)
		for _, roleId := range roleList {
			menus := NewRoleMenuService().GetRoleMenuIdList(roleId)
			menuList = append(menuList, menus...)
		}
		loginData.MenuList = tool.UniqueArrStr(menuList)
	}

	//获取组织架构
	structs := NewAdminStructService().GetStructIdListByAdmin(admin.Id)
	if len(structs) > 0 {
		structInfo := NewStructService().GetInfo(structs[0])
		if structInfo == nil || structInfo.Status != "1" {
			return errors.New("无组织，登录失败")
		}
		loginData.StructId = structInfo.Id
	}

	cookie, err1 := CreateLoginCookie(loginData, maxAge)
	if err1 != nil {
		return errors.New("生成token失败:" + err1.Error())
	}

	ctx.SetCookie("b5gocmf_login", cookie, maxAge, "/", "", false, true)
	return nil
}

func (l *LoginForm) Check() (*AdminModel, error) {
	if l.CaptchaId != "" {
		valid := tool.NewCaptcha("local").VerifyString(l.CaptchaId, l.CaptchaImg)
		if !valid {
			return nil, errors.New("验证码错误")
		}
	}
	if l.UserName == "" {
		return nil, errors.New("账号不能为空")
	}
	model, err := core.NewDao(NewAdminModel()).GetInfoByField("username", l.UserName, "")
	if err != nil {
		return nil, err
	}
	info := model.(*AdminModel)
	if info.Id == "" {
		return nil, errors.New("账号或密码错误")
	}
	s := md5.Sum([]byte(l.Password))
	p := fmt.Sprintf("%x", s)
	if p != info.Password {
		return nil, errors.New("账号或密码错误")
	}
	return info, nil
}
