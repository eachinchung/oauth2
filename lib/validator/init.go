// @Title        init
// @Description
// @Author       Eachin
// @Date         2021/4/10 12:29 上午

package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator
var Validate = validator.New()

// Init 初始化翻译器
func Init() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New()
		uni := ut.New(zhT)
		Trans, _ = uni.GetTranslator("zh")
		err = zhTranslations.RegisterDefaultTranslations(v, Trans)
	}
	return
}

// RemoveTopStruct 去除提示信息中的结构体名称
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
