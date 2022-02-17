package controllers

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/guregu/null.v4"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 全局翻译器
var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 自定义Gin框架中的Validator引擎
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册方法：获取结构体标签对应的值
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("form"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 注册方法：支持验证空值
		v.RegisterCustomTypeFunc(validateNullValuer, null.Int{}, null.String{}, null.Bool{}, null.Time{}, null.Float{})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 支持的语言环境
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language', 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

// removeTopStruct 去除提示信息中结构体的名称
func removeTopStruct(fields map[string]string) map[string]string {
	result := make(map[string]string)
	for field, err := range fields {
		key := field[strings.Index(field, ".")+1:]
		result[key] = err
	}
	return result
}

// validateNullValuer 验证空值
func validateNullValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
	}
	return nil
}
