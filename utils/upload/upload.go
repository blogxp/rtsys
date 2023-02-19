// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

///////////////  文件上传前处理程序及多平台工厂  /////////////////////

package upload

import (
	"errors"
	"mime/multipart"
	"path"
	"rtsys/utils/core"
	"rtsys/utils/tool"
	"rtsys/utils/trans"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	fileExt  = []string{}
	imgExt   = []string{"jpg", "jpeg", "gif", "png", "bmp"} //默认的图片后缀
	videoExt = []string{"mp4", "m3u8", "ogv", "webm"}       //默认的视频后缀
)

// B5Oss  统一对象存储接口   //////////////////////////////////////////////////////////////////////////
type B5Oss interface {
	Upload(file *UploaderFileInfo) (string, error)
}

// NewOss 对象存储实例化
func NewOss() B5Oss {
	switch core.G_CONFIG.Server.OssType {
	case "local":
		return &LocalOss{}
	default:
		return &LocalOss{}
	}
}

// UploaderAction 上传Oss之前的判断过滤  //////////////////////////////////////////////////////////////////
type UploaderAction struct {
	InputName string   //上传表单文件名
	FileType  string   //上传类型 file img video
	Cat       string   //保存分组
	LimitExt  []string //后缀限制
	LimitSize int64    //大小限制
	Width     int      //裁剪宽度 FileType为img时生效
	Height    int      //裁剪高度 FileType为img时生效
}

// UploaderFileInfo 用于上传传递的信息
type UploaderFileInfo struct {
	File   *multipart.FileHeader
	Type   string //文件类型
	Cat    string //分组
	Width  int    //裁剪宽度 FileType为img时生效
	Height int    //裁剪高度 FileType为img时生效
}

// UploaderResult 上传完成结果
type UploaderResult struct {
	Path       string `json:"path"`
	Url        string `json:"url"`
	OriginName string `json:"origin_name"`
	Ext        string `json:"ext"`
}

// NewUploaderAction 实例上传结构体
func NewUploaderAction(types string, args ...UploaderActionWith) (*UploaderAction, error) {
	u := &UploaderAction{FileType: types}
	switch types {
	case "img":
		u.Width = 1920         //默认存储最大宽度为1920 超出等比例缩放
		u.LimitSize = 10 << 20 //图片默认最大10M
		break
	case "video":
		u.LimitSize = 100 << 20 //视频默认最大100M
		break
	case "file":
		u.LimitSize = 200 << 20 //文件默认最大200M
		break
	default:
		return nil, u.error("上传类型错误")
	}
	if len(args) > 0 {
		for _, t := range args {
			t(u)
		}
	}
	return u, nil
}

func (u *UploaderAction) Upload(ctx *gin.Context) (*UploaderResult, error) {
	//读取文件
	info, errRead := u.reader(ctx)
	if errRead != nil {
		return nil, errRead
	}
	if err := u.checkExt(info); err != nil {
		return nil, err
	}
	if err := u.checkSize(info); err != nil {
		return nil, err
	}
	filePath, errUpload := NewOss().Upload(info)
	if errUpload != nil {
		return nil, errUpload
	}
	result := &UploaderResult{
		Path:       filePath,
		Url:        tool.UrlDomain(filePath, true),
		OriginName: info.File.Filename,
		Ext:        strings.ToLower(path.Ext(info.File.Filename))[1:],
	}
	return result, nil
}

// reader 读取上传文件信息
func (u *UploaderAction) reader(ctx *gin.Context) (*UploaderFileInfo, error) {
	if u.InputName == "" {
		u.InputName = "file"
	}
	file, err := ctx.FormFile(u.InputName)
	if err != nil {
		return nil, err
	}
	info := &UploaderFileInfo{File: file, Cat: u.Cat}
	if u.FileType == "img" {
		info.Type = "img"
		info.Width = u.Width
		info.Height = u.Height
	}
	return info, nil
}

// checkExt 检测后缀
func (u *UploaderAction) checkExt(info *UploaderFileInfo) error {
	var extList []string
	switch u.FileType {
	case "file":
		extList = fileExt
		break
	case "img":
		extList = imgExt
		break
	case "video":
		extList = videoExt
		break
	default:
		return u.error("上传类型错误")
	}
	if extList == nil || len(extList) == 0 {
		return nil
	}

	ext := strings.ToLower(path.Ext(info.File.Filename))[1:]
	if !tool.InArray(ext, extList) {
		return u.error("文件类型不满足" + strings.Join(extList, ","))
	}
	return nil
}

// checkSize 文件大小检测
func (u *UploaderAction) checkSize(info *UploaderFileInfo) error {
	if u.LimitSize <= 0 {
		return nil
	}
	if info.File.Size > u.LimitSize {
		return u.error("文件大小超出最大限制：" + trans.TranFileSize(u.LimitSize))
	}
	return nil
}

func (u *UploaderAction) error(msg string) error {
	return errors.New(msg)
}

type UploaderActionWith func(*UploaderAction)

func UploaderWithCat(cat string) UploaderActionWith {
	return func(u *UploaderAction) {
		u.Cat = cat
	}
}
func UploaderWithInput(input string) UploaderActionWith {
	return func(u *UploaderAction) {
		u.InputName = input
	}
}
func UploaderWithSize(size int64) UploaderActionWith {
	return func(u *UploaderAction) {
		u.LimitSize = size
	}
}
func UploaderWithExt(ext []string) UploaderActionWith {
	return func(u *UploaderAction) {
		u.LimitExt = ext
	}
}
func UploaderWithWidthHeight(width int, height int) UploaderActionWith {
	return func(u *UploaderAction) {
		u.Width = width
		u.Height = height
	}
}
