package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func checkPhone(f validator.FieldLevel) bool {
	field := f.Field()
	str := field.String()
	//result, _ := regexp.MatchString(`^1(?:70\d|(?:9[89]|8[0-24-9]|7[135-8]|66|5[0-35-9])\d|3(?:4[0-8]|[0-35-9]\d))\d{7}$`, str)
	result, _ := regexp.MatchString(`^[1](([3][0-9])|([4][0,1,4-9])|([5][0-3,5-9])|([6][2,5,6,7])|([7][0-8])|([8][0-9])|([9][0-3,5-9]))[0-9]{8}$`, str)
	return result
}
func GetError(data interface{}) (errBool bool, errMsg string) {
	validate := validator.New()
	err := validate.RegisterValidation("checkPhone", checkPhone)
	if err != nil {
		fmt.Println("注册失败！")
		return
	}
	errValidate := validate.Struct(data)
	if errValidate != nil {
		for _, err := range errValidate.(validator.ValidationErrors) {
			if err != nil {
				switch err.Tag() {
				case "required":
					return false, err.Field() + " 是必须的"
				case "len":
					return false, err.Field() + " 长度为 " + err.Param()
				case "max":
					return false, err.Field() + " 最大长度长度为 " + err.Param()
				case "min":
					return false, err.Field() + "最小长度为 " + err.Param()
				case "gt":
					return false, err.Field() + " 必须要大于 " + err.Param()
				case "eq":
					return false, err.Field() + " 必须要等于 " + err.Param()
				case "gte":
					return false, err.Field() + " 必须要大于等于 " + err.Param()
				case "lt":
					return false, err.Field() + " 必须要小于 " + err.Param()
				case "lte":
					return false, err.Field() + " 必须要小于等于 " + err.Param()
				case "checkPhone":
					return false, err.Field() + " 格式错误 " + err.Param()
				}
			}
		}
	}
	return true, ""
}
