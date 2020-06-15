package controllers

import (
	"github.com/astaxie/beego"
	"offergo/lib"
)

type baseController struct {
	beego.Controller
}

//response:success
func (b *baseController) responseSuccess(data interface{}) {
	var success lib.Response
	response := success.Success(data)
	b.Data["json"] = &response
	b.ServeJSON()
	b.StopRun()
}

//response:error
func (b *baseController) responseError(msg string) {
	var error lib.Response
	response := error.Error(msg)
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
