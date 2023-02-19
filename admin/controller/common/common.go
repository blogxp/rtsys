package common

import (
	"crypto/md5"
	"fmt"
	"rtsys/admin/lib"
	"rtsys/admin/services"
	system2 "rtsys/common/models/system"
	"rtsys/common/services/system"
	"rtsys/utils/core"
	"rtsys/utils/upload"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CommonsController struct {
	lib.Controller
}

func (c *CommonsController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("repass", false, c.RePass))
	group.POST(c.Dispatch("repass", false, c.RePassSave))
	group.GET(c.Dispatch("cropper", false, c.Cropper))
	group.POST(c.Dispatch("upload_img", false, c.UploadImg))
	group.POST(c.Dispatch("upload_file", false, c.UploadFile))
	group.POST(c.Dispatch("upload_video", false, c.UploadVideo))
}

func NewCommonController() *CommonsController {
	c := &CommonsController{}
	c.Id = "common"
	return c
}
func (c *CommonsController) RePass(ctx *gin.Context) {
	c.Render(ctx, "repass", nil)
}

type rePassStruct struct {
	OldPass     string `json:"old_pass" validate:"required" label:"旧密码"`
	NewPass     string `json:"new_pass" validate:"required,max=20,min=6" label:"新密码"`
	ConfirmPass string `json:"confirm_pass" validate:"required,max=20,min=6" label:"确认密码"`
}

func (c *CommonsController) RePassSave(ctx *gin.Context) {
	params := &rePassStruct{}
	_ = ctx.ShouldBindBodyWith(params, binding.JSON)

	if err := core.G_Validate.Struct(params); err != nil {
		c.Error(ctx, core.G_Validate.GetError(err))
		return
	}
	if params.NewPass != params.ConfirmPass {
		c.Error(ctx, "新密码和确认密码不一致")
		return
	}
	if params.NewPass == params.OldPass {
		c.Error(ctx, "新密码和旧密码不能一致")
		return
	}
	loginInfo := services.GetLoginByCtx(ctx)
	if loginInfo == nil {
		c.Error(ctx, "登录信息获取失败")
		return
	}
	userInfo := system.NewAdminService().GetInfo(loginInfo.Id)
	if userInfo == nil {
		c.Error(ctx, "用户信息获取失败")
		return
	}

	old := fmt.Sprintf("%x", md5.Sum([]byte(params.OldPass)))

	if old != userInfo.Password {
		c.Error(ctx, "旧密码不正确")
		return
	}
	pwd := fmt.Sprintf("%x", md5.Sum([]byte(params.NewPass)))
	_, err := core.NewDao(system2.NewAdminModel()).UpdateNamed(map[string]any{"password": pwd}, map[string]any{"id": userInfo.Id})
	if err != nil {
		c.Error(ctx, "操作失败"+err.Error())
	} else {
		c.Success(ctx, "更新成功")
	}
}

func (c *CommonsController) Cropper(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	cat := ctx.DefaultQuery("cat", "")
	c.Render(ctx, "cropper", gin.H{"id": id, "cat": cat})
}

// UploadImg 上传图片
func (c *CommonsController) UploadImg(ctx *gin.Context) {
	//获取参数
	catParams := ctx.DefaultPostForm("cat", "")
	widthParams := ctx.DefaultPostForm("width", "0")
	heightParams := ctx.DefaultPostForm("height", "0")
	width, _ := strconv.Atoi(widthParams)
	height, _ := strconv.Atoi(heightParams)

	//实例化上传方法
	action, errAction := upload.NewUploaderAction("img", upload.UploaderWithCat(catParams), upload.UploaderWithWidthHeight(width, height))
	if errAction != nil {
		c.Error(ctx, "上传失败："+errAction.Error())
		return
	}
	//上传
	result, errUpload := action.Upload(ctx)
	if errUpload != nil {
		c.Error(ctx, errUpload.Error())
		return
	}
	c.Success(ctx, "上传成功", result)
}

// UploadFile 上传文件
func (c *CommonsController) UploadFile(ctx *gin.Context) {
	//获取参数
	catParams := ctx.DefaultPostForm("cat", "")

	//实例化上传方法
	action, errAction := upload.NewUploaderAction("file", upload.UploaderWithCat(catParams))
	if errAction != nil {
		c.Error(ctx, "上传失败："+errAction.Error())
		return
	}
	//上传
	result, errUpload := action.Upload(ctx)
	if errUpload != nil {
		c.Error(ctx, errUpload.Error())
		return
	}
	c.Success(ctx, "上传成功", result)
}

// UploadVideo 上传视频
func (c *CommonsController) UploadVideo(ctx *gin.Context) {
	//获取参数
	catParams := ctx.DefaultPostForm("cat", "")

	//实例化上传方法
	action, errAction := upload.NewUploaderAction("video", upload.UploaderWithCat(catParams))
	if errAction != nil {
		c.Error(ctx, "上传失败："+errAction.Error())
		return
	}
	//上传
	result, errUpload := action.Upload(ctx)
	if errUpload != nil {
		c.Error(ctx, errUpload.Error())
		return
	}
	c.Success(ctx, "上传成功", result)
}
