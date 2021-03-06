package controllers

import (
	"encoding/json"
	
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//响应json
func (this *BaseController) ResponseJson(isSuccess bool, msg string, data ...interface{}) {
	status := 0
	if isSuccess {
		status = 1
	}
	ret := map[string]interface{}{"status": status, "msg": msg}
	if len(data) > 0 {
		ret["data"] = data[0]
	}
	//this.Data["json"] = ret
	this.ServeJSON()
	this.StopRun()
}

// response info to client
type ResponseInfo struct {
	Status	int			`json:"status"`
	Msg		string		`json:"msg"`
	Other	string		`json:"extra"`
}

// encoding response info to json format
func JsonInfo(info ResponseInfo) []byte {
	result, err := json.Marshal(info)
	if err != nil {
		panic(err.Error())
	}
	return result
}
