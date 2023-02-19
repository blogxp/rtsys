// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

////////////     本地文件上传代码实现       //////////////

package upload

import (
	"crypto/md5"
	"errors"
	"fmt"
	"image"
	"io"
	"math/rand"
	"os"
	"path"
	"rtsys/config"
	"rtsys/utils/core"
	"rtsys/utils/tool"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
)

// LocalOss 本地文件上传处理
type LocalOss struct {
	File      *UploaderFileInfo
	fileName  string
	config    config.OssLocalConf
	storePath string //存储路径
	callPath  string //显示路径
	fullPath  string //最后生成用户展示的全路径+文件名
}

func (u *LocalOss) Upload(file *UploaderFileInfo) (string, error) {
	u.File = file
	u.config = core.G_CONFIG.Oss.Local
	u.getSaveName()
	if err := u.getSavePath(); err != nil {
		return "", err
	}
	if err := u.saveFile(); err != nil {
		return "", err
	}
	return u.fullPath, nil
}

// 保存文件
func (u *LocalOss) saveFile() error {
	// 读取文件
	f, openError := u.File.File.Open()
	if openError != nil {
		return errors.New("读取文件失败" + openError.Error())
	}
	defer f.Close()

	ext := path.Ext(u.File.File.Filename)
	savePath := u.storePath + u.fileName + ext
	u.fullPath = u.callPath + u.fileName + ext

	//图片是否需要重新调整大小
	if u.File.Type == "img" && (u.File.Width > 0 || u.File.Height > 0) {
		imgSrc, errDecode := imaging.Decode(f)
		if errDecode != nil {
			goto Origin
		}
		var imgDsc *image.NRGBA
		if u.File.Width > 0 && u.File.Height > 0 {
			imgDsc = imaging.Fill(imgSrc, u.File.Width, u.File.Height, imaging.Center, imaging.Lanczos)
		} else {
			imgDsc = imaging.Resize(imgSrc, u.File.Width, u.File.Height, imaging.Lanczos)
		}
		err := imaging.Save(imgDsc, savePath)
		if err != nil {
			goto Origin
		}
		return nil
	}

Origin:
	//创建保存文件
	out, createErr := os.Create(savePath)
	if createErr != nil {
		return errors.New("创建文件失败" + createErr.Error())
	}
	defer out.Close()

	//保存文件
	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		return errors.New("保存文件失败:" + copyErr.Error())
	}

	return nil
}

// 获取存储路径和展示路径
func (u *LocalOss) getSavePath() error {
	store := u.config.StorePath
	call := u.config.CallPath

	randPath := time.Now().Format("2006/01/02/")
	cat := u.File.Cat
	if cat == "" {
		cat = u.File.Type
	}
	u.storePath = store + cat + "/" + randPath
	u.callPath = call + cat + "/" + randPath

	//创建文件夹
	if err := tool.DirCreate(u.storePath); err != nil {
		return errors.New("保存目录创建失败" + err.Error())
	}
	return nil
}

// 生成文件名称
func (u *LocalOss) getSaveName() {
	timestamp := time.Now().UnixNano()
	rand.Seed(timestamp)
	random := rand.Int63n(1000)
	name := u.File.File.Filename + strconv.FormatInt(timestamp, 10) + strconv.FormatInt(random, 10)
	u.fileName = fmt.Sprintf("%x", md5.Sum([]byte(name)))
}
