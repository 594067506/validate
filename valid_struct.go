package validate

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)


type validStruct struct {}

func New() *validStruct  {
	return &validStruct{}
}

//基于 github.com/go-playground/validator 实现 struct 的数据验证，错误信息支持返回中文
// in_params 是一个结构体的指针，语法遵循validator规则
func (v * validStruct)Validate( in_params interface{} )   (err error) {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name:=fld.Tag.Get("json")
		return name
	})

	if err=zh_translations.RegisterDefaultTranslations(validate, trans);err!=nil {
		return
	}

	if err = validate.Struct(in_params); err != nil {
		for _,ierr:=range err.(validator.ValidationErrors){
			err = errors.New(ierr.Translate(trans))
			return
		}
	}
	return
}