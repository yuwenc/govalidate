package main

import (
	"fmt"
	"reflect"
	"time"
	"validate"
)

func main() {
	v := validate.New()

	v.UseFormat("date", func(f *validate.Field) bool {
		switch f.Kind {
		case reflect.String:
			if _, err := time.Parse("2006-01-02", f.Val.String()); err == nil {
				return true
			}
		}
		return false
	})

	data := struct {
		Account   string `validate:"format=email > 邮箱格式错误"`
		Name      string `validate:"empty=true | format=trim_space & gt=4 > 字符必须大于4个"`
		Age       int    `validate:"o_interval=10,100 > 年龄需要大于10小于100"`
		Mobile    string `validate:"format=cn_mobile > 手机格式错误"`
		Status    int    `validate:"in=0,1 >状态值错误"`
		DateStart string `validate:"format=date>日期格式错误"`
	}{
		Account:   "even@qq.com",
		Name:      "eventt ",
		Age:       6,
		Mobile:    "1361173787",
		Status:    -1,
		DateStart: "2022-05",
	}
	if !v.Struct(&data).Check() {
		for _, val := range v.GetErrors() {
			fmt.Println(val.Msg)
		}
	}

	fmt.Printf("%+v \n", data)
}
