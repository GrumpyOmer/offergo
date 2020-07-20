package controllers

import (
	"github.com/astaxie/beego"
)


//响应结构体
type result struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

type baseController struct {
	beego.Controller
}

//response:success
func (b *baseController) responseSuccess(data interface{}) {
	response := b.success(data)
	b.Data["json"] = &response
	b.ServeJSON()
	b.StopRun()
}

//response:error
func (b *baseController) responseError(msg string) {
	response := b.error(msg)
	b.Data["json"] = &response
	b.ServeJSON()
	b.StopRun()
}

//request:filter
func (b *baseController) requestFilter(data map[string]interface{}) bool {
	var errString string
	for k, v := range data {
		if v == "" {
			if errString == "" {
				errString += k + " "
			} else {
				errString += "、" + k + " "
			}
		}
	}
	if errString != "" {
		errString += "not define"
		b.responseError(errString)
		return false
	}
	return true
}

//result:success
func (b *baseController) success(data interface{}) result {
	result := result{
		Code: 200,
		Msg:  "ok",
		Data: data,
	}
	return result
}

//result:error
func (b *baseController) error(msg string) result {
	result := result{
		Code: 400,
		Msg:  msg,
		Data: nil,
	}
	return result
}
