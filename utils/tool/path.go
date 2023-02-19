// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package tool

import (
	"os"
	"path/filepath"
	"rtsys/utils/core"
	"strings"
)

// DirCreate 创建文件夹
func DirCreate(path string) error {
	if strings.Index(path, "/") == 0 || strings.Index(path, "\\") == 0 {
		path = "." + path
	}
	if err := os.MkdirAll(filepath.Clean(path), os.ModePerm); err != nil {
		return err
	}
	return nil
}

// UrlDomain 拼接域名
func UrlDomain(path string, isFile bool) string {
	if path == "" {
		return path
	}
	if strings.Index(path, "://") > -1 {
		return path
	}
	domain := core.G_CONFIG.Server.Domain
	port := core.G_CONFIG.Server.Port
	if port != "80" {
		domain = domain + ":" + port
	}
	if isFile && core.G_CONFIG.Server.OssDomain != "" {
		domain = core.G_CONFIG.Server.OssDomain
	}

	return domain + "/" + strings.TrimLeft(path, "/")
}

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	file, err := os.Stat(path)
	if err == nil {
		return !file.IsDir()
	} else {
		return !os.IsNotExist(err)
	}
}
