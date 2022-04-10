package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en" //多语言包
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en" //validator 的翻译器
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

//国际化处理 -- 编写针对 validator 的语言包翻译的相关功能

//原因:go-playground/validator 默认的错误信息是英文.

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale") //希望的语言
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			v.RegisterTagNameFunc(func(fld reflect.StructField) string { //实现返回的错误为form标签内容
				name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
				return name
			})
			//v.RegisterStructValidation(models.SignUpParamStructLevelValidation, models.SignUpParam{}, models.ListPostProParam{})
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}

/*
在自定义中间件 Translations 中，我们针对 i18n 利用了第三方开源库去实现这块功能，分别如下：

go-playground/locales：多语言包，是从 CLDR 项目（Unicode 通用语言环境数据存储库）生成的一组多语言环境，主要在 i18n 软件包中使用，该库是与 universal-translator 配套使用的。
go-playground/universal-translator：通用翻译器，是一个使用 CLDR 数据 + 复数规则的 Go 语言 i18n 转换器。
go-playground/validator/v10/translations：validator 的翻译器。
*/
