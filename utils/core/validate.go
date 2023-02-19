// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package core

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

type ValidateCtx struct {
	*validator.Validate
	trans ut.Translator
}

func LoadValidator() {
	zhTranslator := zh.New()
	uni := ut.New(zhTranslator, zhTranslator)

	trans, _ := uni.GetTranslator("zh")

	validate := validator.New()
	//通过自定义标签label来替换字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)

	G_Validate = &ValidateCtx{validate, trans}
}

func (vc *ValidateCtx) GetError(errs error) string {
	errStr := ""
	for _, err := range errs.(validator.ValidationErrors) {
		if vc.trans != nil {
			errStr = err.Translate(vc.trans)
		} else {
			errStr = err.Field() + "验证不符合" + err.Tag()
		}
		break
	}
	return errStr
}
