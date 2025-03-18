package res

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// ValidateErrors 封装验证错误处理
func ValidateErrors(err error, req interface{}) string {
	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		if f, ok := reflect.TypeOf(req).Elem().FieldByName(field); ok {
			// 读取字段的 json 标签名称（用于返回更友好的字段名）
			jsonTag := f.Tag.Get("json")
			fieldName := jsonTag
			if fieldName == "" {
				fieldName = field
			}

			// 根据验证规则返回错误信息
			switch e.Tag() {
			case "required":
				return fieldName + " 字段不能为空"
			case "min":
				return fieldName + " 字段值不能小于 " + e.Param()
			case "max":
				return fieldName + " 字段值不能大于 " + e.Param()
			case "email":
				return fieldName + " 字段必须是有效的邮箱地址"
			default:
				return fieldName + " 字段验证失败"
			}
		}
	}
	return ""
}
